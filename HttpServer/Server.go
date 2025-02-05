package HttpServer

import (
	"fmt"
	"net/http"
)

// 处理 "/" 路径的请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "text/plain")
	// 向客户端响应数据
	fmt.Fprintf(w, "Hello, World!")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// 只处理 POST 请求
	if r.Method == http.MethodPost {
		// 解析表单数据
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		// 获取表单字段
		name := r.FormValue("name")
		age := r.FormValue("age")

		// 输出接收到的表单数据
		fmt.Fprintf(w, "Received Name: %s, Age: %s", name, age)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func StartServer() {
	// 注册路由
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/post", PostHandler)

	// 启动 HTTP 服务，监听 8080 端口
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
