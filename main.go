package main

import (
	"fmt"
	"github.com/ricejson/example-common/model"
	"github.com/ricejson/example-common/service"
	"github.com/ricejson/example-consumer/impl"
)

func main() {
	var userService service.UserService = impl.NewUserServiceImpl()
	var user = model.User{}
	user.SetName("ricejson")
	getUser, err := userService.GetUser()
	if err != nil {
		fmt.Println("user == nil")
	} else {
		fmt.Println(getUser.GetName())
	}
}
