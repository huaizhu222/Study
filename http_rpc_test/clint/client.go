package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 请求 URL
	url := "http://localhost:1234/jsonrpc"

	// 请求 Body (JSON 数据)
	requestBody := map[string]interface{}{
		"id":     0,
		"params": []interface{}{"bobby"}, // 修正 params 为数组
		"method": "HelloService.Hello",
	}

	// 将 Go 对象编码为 JSON
	jsonData, _ := json.Marshal(requestBody)

	// 创建 HTTP 请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 发送请求
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Response:", string(body))
}
