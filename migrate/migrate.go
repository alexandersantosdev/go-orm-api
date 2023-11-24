package main

import (
	"github.com/alexandersantosdev/go-orm-api/initializers"
	"github.com/alexandersantosdev/go-orm-api/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
