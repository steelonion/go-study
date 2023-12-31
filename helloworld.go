package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// 使用互斥锁确保并发安全
var (
	count   int
	countMu sync.Mutex
)

// Message 结构定义了要返回的JSON消息格式
type Message struct {
	Text  string `json:"text"`
	Count int    `json:"int"`
}

func main() {
	var rp RollPoint
	rp = FooRandom{}
	fmt.Println(rp.GetPoint())
	rp = FooRandom2{}
	fmt.Println(rp.GetPoint())

	// 设置处理函数
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/addvarb", addVarbHandler)

	// 启动HTTP服务器并监听端口
	http.ListenAndServe(":8080", nil)
}

func addVarbHandler(w http.ResponseWriter, r *http.Request) {

	// 使用互斥锁保护计数器的并发访问
	countMu.Lock()
	defer countMu.Unlock()

	switch r.Method {
	case http.MethodPost:

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		str := r.Form.Get("name")
		countstr := r.Form.Get("count")

		paramInt, err := strconv.Atoi(countstr)
		if err != nil {
			http.Error(w, "Failed to convert parameter to int", http.StatusBadRequest)
			return
		}

		// 添加新的varb
		testvarbList = append(testvarbList, testvarb{StringField: str, IntField: paramInt})

		// 将消息转换为JSON格式
		jsonData, err := json.Marshal(testvarbList)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// 设置响应头为JSON格式
		w.Header().Set("Content-Type", "application/json")

		// 发送JSON响应
		w.Write(jsonData)
		break
	default:
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func countHandler(w http.ResponseWriter, r *http.Request) {
	// 使用互斥锁保护计数器的并发访问
	countMu.Lock()
	defer countMu.Unlock()

	// 每次调用count路径时，增加计数值
	count++

	// 创建要返回的消息
	message := Message{Text: "Count!", Count: count}

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
