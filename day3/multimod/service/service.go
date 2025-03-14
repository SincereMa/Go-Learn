package service

import (
	"fmt"
	"multimod/api"
)

func Server(name string) {
	fmt.Println(api.Greet(name))
}
