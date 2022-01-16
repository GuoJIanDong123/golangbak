package main

import (
	"fmt"
	"net/http"
)

func DisplayHeadersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method:%s URL:% Protocol:%s \n", r.Method, r.URL, r.Proto)
	//遍历所有请求
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q,Value %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr=%q\n", r.RemoteAddr)
	//通过key获取指定请求头的值
	fmt.Fprintf(w, "\n\nFinding value of \"Accept\"%q", r.Header["Accept"])
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// 一个非常简单的健康检查实现：如果此 HTTP 接口调用成功，则表示应用健康
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/display_headers", DisplayHeadersHandler)
	http.HandleFunc("/healthz", HealthCheckHandler)
	http.ListenAndServe(":8000", nil)

}
