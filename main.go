package main

import (
	"CI-CD_test/pkg"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	pkg.Handler(r)
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to load env")
	}
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal("unable to start on port")
	}
}
