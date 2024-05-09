package router

import (
	"idp/configs"
	"idp/internal/database"
	redisclient "idp/internal/database/redis_client"
	user_handler "idp/internal/handler/user_handler"
	user_repo "idp/internal/repository/userrepository"
	"idp/internal/service/userservice"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, cfg configs.Config) {
	dbConnect := database.CreateConnection(cfg)
	redisConn := database.CreateRedisConnection(cfg)
	database.RunMigration(dbConnect)

	//emailservice := emailservice.NewEmailService(cfg.Email)
	userRepo := user_repo.NewUserRepository(dbConnect)
	userService := userservice.NewUserService(userRepo)
	redisClt := redisclient.NewRedisClient(redisConn)
	userHandler := user_handler.NewUserHandler(userService, redisClt)

	group := engine.Group("user/api")
	{
		//User apis
		group.POST("/register", userHandler.Register)
		group.POST("/login", userHandler.Login)
		group.POST("/logout", userHandler.Logout)
	}

}
