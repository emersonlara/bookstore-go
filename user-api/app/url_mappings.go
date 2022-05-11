package app

import (
	"user-api/controllers/ping"
	"user-api/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)

}
