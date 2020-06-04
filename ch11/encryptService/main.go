package main

import (
	"fmt"

	proto "github.com/YoungsoonLee/restful_go/ch11/encryptService/proto"
	micro "github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("encrypter"),
	)

	service.Init()

	proto.RegisterEncrypterHandler(service.Server(), new(Encrypter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
