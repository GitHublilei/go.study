package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:9080/demo/?debug=true")
	// if err != nil {
	// 	fmt.Printf("get failed, err:%v\n", err)
	// 	return
	// }

	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9080/demo/")
	data.Set("name", "悟空")
	data.Set("age", "2000")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Printf("request get failed, err:%v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("do req failed, err:%v\n", err)
		return
	}

	defer resp.Body.Close()

	// 从resp中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
