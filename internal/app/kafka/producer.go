package kafka

import (
	"fmt"
	"math/rand"
	"os"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/alexsalniy/test-service/internal/app/kafka/utils"
	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
)

var names = [...]string{"Jane", "Joe", "Alex", "Haley", "Jinny", "Zane", "Ivan", "Adam", "Eytan", "Goldie", "Hanna", "Naomi", "Therese", "Annette", "Sasha", "Sveta", "Saveliy", "Timur", "Artur", "Josefine", "Marlen"}
var surnames = [...]string{"Dubina", "Kolchak", "Mendeleev", "Korolev", "Gurin", "Medvedev", "Bertrand", "Lombard", "Lefort", "Steiner", "Marx", "Schumann", "Stark", "Fisher", "Adams", "Johnson", "Boyle"}
var patronymics = [...]string{"Sergeevich", "Mihalich", "Alexandrovich", "Dmitrievich", "Alexeevich", "Ivanich", "Fedorovich", "Genadievich"}

func Producer() {
	
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n", os.Args[0])
		os.Exit(1)	
	} 
	configFile := os.Args[1]
	conf := utils.ReadConfig(configFile)

	topic := "purchases"
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	for n := 0; n < 5; n++ {
		name := names[rand.Intn(len(names))]
		surname := surnames[rand.Intn(len(surnames))]
		patronymic := patronymics[rand.Intn(len(patronymics))]
		fio := model.FIO{}
		switch n % 2 {
			case 0:
				fio = model.FIO{
					Name: name,
					Surname: surname,
					Patronymic: "",
				}
			default: 
				fio = model.FIO{
					Name: name,
					Surname: surname,
					Patronymic: patronymic,
				}
		}
		jsonFio, err := json.Marshal(fio)
  	if err != nil {
    	fmt.Println(err)
			return
		}
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key: 			[]byte(name),
			Value: 		[]byte(jsonFio),	
		}, nil)
	}

	p.Flush(15 * 1000)
	p.Close()
}

func ReadConfig(configFile string) {
	panic("unimplemented")
}