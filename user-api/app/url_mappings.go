package app

import (
	"user-api/controllers/ping"
	"user-api/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)

}
