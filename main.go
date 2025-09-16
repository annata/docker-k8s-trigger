package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
)

var secret = "6FGS959U9epLIJ6crvu4l2TsdpD4Ozyz2M7JEravear"
var logger = log.New(os.Stdout, "[k8s_trigger]", log.Ldate|log.Ltime)

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
	if len(p) != 4 && len(p) != 6 {
		handleResponse(response, errors.New("参数错误"))
		return
	}
	if p[1] != secret {
		handleResponse(response, errors.New("参数错误"))
		return
	}
	if p[2] == "kube-system" {
		handleResponse(response, errors.New("参数错误"))
		return
	}
	if len(p) == 6 {
		logger.Printf("triggerVersion,namespace: %s, name: %s, container: %s, tag: %s", p[2], p[3], p[4], p[5])
		err := triggerVersion(p[2], p[3], p[4], p[5])
		handleResponse(response, err)
	} else {
		logger.Printf("trigger,namespace: %s, name: %s", p[2], p[3])
		err := trigger(p[2], p[3])
		handleResponse(response, err)
	}
}

func handleResponse(response http.ResponseWriter, err error) {
	response.Header().Add("Content-Type", "application/json")
	if err == nil {
		_, _ = response.Write([]byte("{\"code\":\"0\",\"msg\":\"请求成功\"}"))
	} else {
		_, _ = response.Write([]byte("{\"code\":\"100\",\"msg\":\"" + err.Error() + "\"}"))
	}
}
