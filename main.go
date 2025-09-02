package main

import (
	"pet_marketplace_users/config"
	"pet_marketplace_users/routes"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	config.ConnectDB()
	r := routes.SetupRoutes()
	logrus.Info("Сервер запускается на порту 8080")
	r.Run(":8080")
}
