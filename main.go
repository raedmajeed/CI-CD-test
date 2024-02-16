package main

import (
	"CI-CD_test/pkg"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	r := gin.Default()
	pkg.Handler(r)
	testVal := flag.String("v", "1", "test value")
	flag.Parse()
	log.Printf("Run v %v successful", *testVal)
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to load env")
	}
	if err := r.Run(":8089"); err != nil {
		log.Fatal("unable to start on port")
	}
}
