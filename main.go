package main

import (
	"fmt"
	"go-todo-app/app/models"
)

func main() {
	//fmt.Println(config.Config.Port)
	//fmt.Println(config.Config.SQLDriver)
	//fmt.Println(config.Config.DbName)
	//fmt.Println(config.Config.LogFile)

	//log.Println("test")

	fmt.Println(models.Db)

	//u := models.User{}
	//u.Name = "test"
	//u.Email = "test@example.com"
	//u.PassWord = "testtest"
	//
	//u.CreateUser()

	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "Test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
