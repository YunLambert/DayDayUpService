package router

import (
	"daydayup/db"
	"daydayup/db/dao"
	"daydayup/handler"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Init(sqlite db.SQLite) *gin.Engine {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.BestSpeed))

	helloHost := router.Group("/hello")
	{
		helloHost.GET("", handler.Greetings)
	}

	ds := handler.DailySign{
		Dao: dao.NewDailySignDAO(sqlite),
	}
	signHost := router.Group("/sign")
	{
		signHost.POST("", ds.Sign)
		signHost.POST("summary", ds.GetSummary)
	}

	return router
}
