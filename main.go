package main

import (
	"CI-CD_test/pkg"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	r := gin.Default()
	pkg.Handler(r)
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to load env")
	}
	if err := r.Run(":8089"); err != nil {
		log.Fatal("unable to start on port")
	}
}
