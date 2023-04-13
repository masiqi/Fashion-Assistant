package main

import (
	"context"
	"fmt"
	ernie "github.com/zjy282/ernie-api"
)

func main() {
	ctx := context.Background()
	req := &ernie.OAuthTokenRequest{
		ClientID:     "$$id",
		ClientSecret: "$$secret",
	}
	response, err := ernie.CreateOAuthToken(ctx, req)
	fmt.Println("vim-go")
	fmt.Println(response)
	fmt.Println(err)
}
