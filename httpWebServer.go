package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		fmt.Println("\n=========== 收到请求信息===============\n")

		scheme := "http://"
		if request.TLS != nil {
			scheme = "https://"
		}
		fmt.Println(strings.Join([]string{request.Method, "  ", scheme, request.Host, request.RequestURI}, ""))
		fmt.Println("\n=========== 打印Header列表：===========\n")
		if len(request.Header) > 0 {
			for k, v := range request.Header {
				fmt.Printf("%s=%s\n", k, v[0])
			}
		}
		fmt.Println("\n=========== 打印Form参数列表：===========\n")
		request.ParseForm()
		if len(request.Form) > 0 {
			for k, v := range request.Form {
				fmt.Printf("\n%s=%s\n", k, v[0])
			}
		}
		fmt.Println("\n=========== 请求body体：===========\n")

		b, _ := ioutil.ReadAll(request.Body)
		fmt.Println(string(b))
	})
	http.ListenAndServe(":7000", nil)
}
