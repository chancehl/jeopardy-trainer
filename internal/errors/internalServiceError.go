package errors

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewInternalServiceError(msg string, e error) gin.H {
	return gin.H{"error": fmt.Sprintf("%s: %v", msg, e)}
}
