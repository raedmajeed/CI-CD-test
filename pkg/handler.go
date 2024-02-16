package pkg

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Handler(r *gin.Engine) {
	r.GET("/api/v1/test", testFunction)
}

func testFunction(ctx *gin.Context) {
	log.Println("RUN IS SUCCESSFUL")
}
