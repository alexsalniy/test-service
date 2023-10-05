package apiserver

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"sync"

	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
	"github.com/alexsalniy/test-service/internal/app/kafka/utils"
	"github.com/alexsalniy/test-service/internal/store"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type fioStore struct {
	store store.Store
}

var FioList = make([][]byte, 0)

func Consumer(store store.Store, wg *sync.WaitGroup) {
	defer wg.Done()

	s :=  &fioStore{
		store:	store,
	}

	configFile := "getting-started.properties"
	conf := utils.ReadConfig(configFile)
	conf["group.id"] = "kafka-go-getting-started"
	conf["auto.offset.reset"] = "earliest"

	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	topic := "purchases"
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Some error")
	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)



	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))		

			var newFIO model.ExtendedFIO
			readerr := json.Unmarshal(ev.Value, &newFIO)
			if readerr != nil {
				fmt.Printf("Some error")
			}

			e := &model.ExtendedFIO{
				Name: newFIO.Name,
				Surname: newFIO.Surname,
				Patronymic: newFIO.Patronymic,
			}

			if err := s.store.ExtFIO().Create(e); err != nil {
				fmt.Println(err)
			}
				
		}
	}

c.Close()
}