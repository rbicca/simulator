package main

import (
	"fmt"
	route2 "github.com/rbicca/simulator/application/route"
	"github.com/joho/godotenv"
	"github.com/rbicca/simulator/infra/kafka"
)

func init(){
	err := gotodenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main(){


	producer := kafka.NewKafkaProducer()
	kafka.Publish("Ola Mella", "readTest"producer)

	// route := route2.Route{
	// 	ID: 		"1",
	// 	ClientID:	"1",
	// } 

	// route.LoadPositions()

	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[2]);

}