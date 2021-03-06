package key

import (
	"github.com/gin-gonic/gin"
)

func RouteKey(router *gin.Engine) {
	secretKey := router.Group("/key")
	secretKey.PUT("/intermediateKey", updateIntermediateKeyHandler)
	secretKey.PUT("/openingKey", updateOpeningKeyHandler)
	secretKey.POST("", createKeyHandler)
	secretKey.POST("/generateRsaKey", generateRsaKeyHandler)
}
