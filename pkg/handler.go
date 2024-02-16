package pkg

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

func Handler(r *gin.Engine) {
	r.GET("/api/v1/test", testFunction)
}

func testFunction(ctx *gin.Context) {
	testVal := flag.String("v", "1", "test value")
	flag.Parse()
	log.Printf("Run %v successful", *testVal)
}
