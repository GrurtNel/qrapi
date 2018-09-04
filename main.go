package main

import (
	"log"
	"os"
	// 1 load first
	_ "qrapi/config"
	// 2
	"github.com/gin-gonic/gin"
	"net/http"
	"qrapi/api"
	"qrapi/middleware"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
func main() {
	r := gin.New()
	r.StaticFS("/app", http.Dir("./app")).Use(middleware.AddStaticHeader())
	r.Use(gin.Logger(), middleware.AddHeader(), middleware.RecoveryWithWriter())
	r.StaticFS("/static", http.Dir("./upload"))
	api.NewApiServer(r.Group("api"))
	// Listen and serve on 0.0.0.0:8080
	r.Run(GetPort())
}
