package main

import (
	"encoding/json"
	"net/http"
)

// Message 结构定义了要返回的JSON消息格式
type Message struct {
	Text string `json:"text"`
}

func main() {
	// 设置处理函数
	http.HandleFunc("/hello", helloHandler)

	// 启动HTTP服务器并监听端口
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 创建要返回的消息
	message := Message{Text: "Hello, World!"}

	// 将消息转换为JSON格式
	jsonData, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 设置响应头为JSON格式
	w.Header().Set("Content-Type", "application/json")

	// 发送JSON响应
	w.Write(jsonData)
}

