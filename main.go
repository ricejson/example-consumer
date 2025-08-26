package main

import (
	"fmt"
	"github.com/ricejson/example-common/model"
	"github.com/ricejson/example-common/service"
)

func main() {
	// todo 消费者需要替换为对应实现类对象
	var userService service.UserService
	var user = model.User{}
	user.SetName("ricejson")
	getUser, err := userService.GetUser(user)
	if err != nil {
		fmt.Println("user == nil")
	} else {
		fmt.Println(getUser.GetName())
	}
}
