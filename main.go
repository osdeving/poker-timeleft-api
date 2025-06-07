package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/osdeving/poker-timeleft-api/config"
	"github.com/osdeving/poker-timeleft-api/controllers"
	"github.com/osdeving/poker-timeleft-api/database"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	t := r.Group("/tournaments")
	{
		t.POST("", controllers.CreateTournament)
		t.GET("", controllers.GetAllTournaments)
		t.GET(":id", controllers.GetTournament)
		t.PUT(":id", controllers.UpdateTournament)
		t.DELETE(":id", controllers.DeleteTournament)
	}

	return r
}

func main() {
	config.LoadConfig()
	database.ConnectMongo(config.AppConfig.MongoURI)

	r := setupRouter()
	r.Run(":8080")
}
