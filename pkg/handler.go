package pkg

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Handler(r *gin.Engine) {
	r.GET("/api/v1/test", testFunction)
}

func testFunction(ctx *gin.Context) {
	log.Println("RUN IS SUCCESSFUL here")
	ctx.JSON(http.StatusOK, gin.H{
		"success": "false",
	})
}
