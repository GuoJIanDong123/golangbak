package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	//创建mux.Router路由器示例
	router := mux.NewRouter().StrictSlash(true)

	//应用请求日志中间件
	router.Use(loggingRequestInfo)

	//遍历web.go中定义的所有webRoutes
	for _, route := range webRoutes {
		//将每个web路由应用到路由器
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router

}

//记录请求日志信息中间件
func loggingRequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//打印请求URL明细
		url, _ := json.Marshal(r.URL)
		log.Println(string(url))
		next.ServeHTTP(w, r)
	})
}
