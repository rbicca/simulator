package main

import (
	"fmt"
	//route2 "github.com/rbicca/simulator/application/route"
	kafka2 "github.com/rbicca/simulator/application/kafka"
	"github.com/joho/godotenv"
	"github.com/rbicca/simulator/infra/kafka"
	"log"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main(){

	//** Teste mocado para consumir rotas
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}


	//** Teste mocado para consumir mensagem enviadas
	// msgChan := make(chan *ckafka.Message)
	// consumer := kafka.NewKafkaConsumer(msgChan)
	// go consumer.Consume()

	// for msg := range msgChan {
	// 	fmt.Println(string(msg.Value))
	// }

	//** Teste mocado para enviar mensagem - loopzinho maroto no fim para manter app rodando
	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("Ola Mella", "readTest", producer)

	// for {
	// 	_ = 1
	// }

	//** Teste mocado para carga e exibição de rotas
	// route := route2.Route{
	// 	ID: 		"1",
	// 	ClientID:	"1",
	// } 

	// route.LoadPositions()

	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[2]);

}