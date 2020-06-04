package main

import (
	"context"

	proto "github.com/YoungsoonLee/restful_go/ch11/encryptService/proto"
)

type Encrypter struct{}

func (g *Encrypter) Encrypt(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Result = EncryptString(req.Key, req.Message)
	return nil
}
func (g *Encrypter) Decrypt(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Result = DecryptString(req.Key, req.Message)
	return nil
}
