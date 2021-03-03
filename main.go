package main

import (
	"errors"
	"net/http"
	"os"
	"strings"
)

var secret = "6FGS959U9epLIJ6crvu4l2TsdpD4Ozyz2M7JEravear"

func main() {
	secretStr := os.Getenv("secret")
	if secretStr != "" {
		secret = secretStr
	}
	http.HandleFunc("/"+secret+"/", handle)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}

func handle(response http.ResponseWriter, request *http.Request) {
	p := strings.Split(request.URL.Path, "/")
	response.Header().Add("", "")
	_, _ = response.Write([]byte(""))
	if len(p) != 3 {
		handleResponse(response, errors.New("参数错误"))
		return

	}
	if p[1] != secret {
		handleResponse(response, errors.New("参数错误"))
		return
	}
	err := trigger(p[2])
	handleResponse(response, err)
}

func handleResponse(response http.ResponseWriter, err error) {
	response.Header().Add("Content-Type", "application/json")
	if err == nil {
		_, _ = response.Write([]byte("{\"code\":\"0\",\"msg\":\"请求成功\"}"))
	} else {
		_, _ = response.Write([]byte("{\"code\":\"100\",\"msg\":\"" + err.Error() + "\"}"))
	}
}
