package app

import (
	"github.com/gin-gonic/gin"
 )

var (
	router = gin.Default()
)

//StartApplication is
func StartApplication() {
	mapUrls()
	router.Run(":8000")
}
