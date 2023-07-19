package main

import (
	"ipProject/service"
	"ipProject/config"
	"log"

)

func main(){

	log.Printf("Started the main service!!")
	config.Init()
	service.HandleRequests()

}