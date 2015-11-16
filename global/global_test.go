package global

import (
	"net/http"

	_ "github.com/walf443/go-loghttp/global"
)

var client = &http.Client{}

func Example() {
	// Automatically logs requests/responses
	//
	// For example:
	// 2014/12/01 19:33:10 ---> GET http://www.example.com/
	// 2014/12/01 19:33:11 <--- 200 http://www.example.com/
	client.Get("http://www.example.com/")
}
