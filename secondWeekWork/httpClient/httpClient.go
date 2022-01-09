package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	testHttpPost()
}

func testHttpPost() {
	//构建参数
	data := url.Values{
		"theCityName": {"重庆"},
	}
	//参数转化为body
	reader := strings.NewReader(data.Encode())
	//发起post请求MIME格式
	response, err := http.Post("https://www.qidian.com/", "application/x-www-form-urlencoded", reader)

	CheckErr(err)
	fmt.Printf("响应状态码:%v\n", response.StatusCode)

	if response.StatusCode == 200 {
		//操作相应数据
		defer response.Body.Close()
		fmt.Println("网络请求成功")
		CheckErr(err)
	} else {
		fmt.Println("请求失败", response.Status)
	}

}

func testClientGet() {
	//创建客户端
	client := http.Client{}
	//通过client去请求

	response, err := client.Get("https://www.qidian.com/")
	CheckErr(err)

	fmt.Printf("状态码：%v\n", response.Status)

	if response.StatusCode == 200 {
		fmt.Println("网络请求成功")
		defer response.Body.Close()
	}
}

func testHttpNewRequest() {
	//1.创建一个客户端
	client := http.Client{}
	//2.创建一个请求，请求方式既可以是GET，也可以是POST
	request, err := http.NewRequest("GET", "https://www.qidian.com/", nil)
	CheckErr(err)

	//3.客户端发送请求
	cookName := &http.Cookie{Name: "username", Value: "Steven"}

	//添加cookie
	request.AddCookie(cookName)
	response, err := client.Do(request)
	CheckErr(err)

	//设置请求头
	request.Header.Set("Accept-Lanauage", "zh-cn")
	defer response.Body.Close()

	//查看请求头的数据
	fmt.Printf("Header:%+v\n", request.Header)
	fmt.Printf("响应状态码：%v\n", response.StatusCode)

	//4.操作数据
	if response.StatusCode == 200 {
		//data,err := ioutil.ReadAll(response.Body)
		fmt.Println("网络请求成功")
		CheckErr(err)
	} else {
		fmt.Println("网络请求失败", response.Status)
	}
}

//检查错误
func CheckErr(err error) {

	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现异常：", ins.Error())
		}
	}()

	if err != nil {
		panic(err)
	}
}
