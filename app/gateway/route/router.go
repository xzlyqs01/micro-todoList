package route

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"micro-todoList/micro-todoList/app/gateway/http"
	"micro-todoList/micro-todoList/app/gateway/middleware"

)

func NewRouter()*gin.Engine{
	ginRouter:=gin.Default()
	ginRouter.Use(middleware.Cors())
	store:=cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession",store))
	v1:=ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200,"success")
		})
		//用户服务
		v1.POST("/user/register",)
	}
}







