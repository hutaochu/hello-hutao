package services

import "fmt"

type Hello struct {
	Name    string
	Message string
}

func SayHello(name string) Hello {
	return Hello{
		Name:    name,
		Message: fmt.Sprintf("Hello, %s", name),
	}
}
