package kafka

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/alexsalniy/test-service/internal/app/kafka/utils"
	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
	"github.com/alexsalniy/test-service/internal/store/sqlstore"
)

type ExtFIORepo struct {
	store *sqlstore.ExtFIORepository
}

var FioList = make([][]byte, 0)

func Consumer(s *ExtFIORepo) {
	// func (s *server) consumer() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n", os.Args[0])
		os.Exit(1)
	}

	configFile := os.Args[1]
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

			var newFIO model.FIO
			readerr := json.Unmarshal(ev.Value, &newFIO)
			if readerr != nil {
				fmt.Printf("Some error")
			}

			e := &model.ExtendedFIO{
				Name: newFIO.Name,
				Surname: newFIO.Surname,
				Patronymic: newFIO.Patronymic,
			}

			if err := s.store.Create(e); err != nil {
				fmt.Println(err)
			}
				
		}
	}

c.Close()
}