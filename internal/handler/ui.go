package handler

import "github.com/gin-gonic/gin"

func ServeSPA(ctx *gin.Context) {
	ctx.File("./web/dist/index.html")
}

func HandleSPARoute(ctx *gin.Context) {
	ctx.File("./web/dist/index.html")
}
