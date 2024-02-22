package main

import (
	"fmt"
	"httpserver/controller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/httprequest", controller.HandleRequest)
	router.POST("/channelrequest", controller.ChannalRequest)
	router.POST("/goworker", controller.Worker)
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
