package api

import "github.com/gin-gonic/gin"

// TODO fix me
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func response(data any) gin.H {
	return gin.H{
		"data": data,
	}
}
