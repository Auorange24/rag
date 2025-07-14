package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// GenerateRequest 定义请求结构体
type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

func generate() {
	// 定义请求数据
	requestData := GenerateRequest{
		Model:  "deepseek-r1:1.5b",
		Prompt: "1+1",
		Stream: false,
	}

	// 将请求数据转换为JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("JSON编码失败: %v\n", err)
		os.Exit(1)
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", "http://127.0.0.1:11434/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		os.Exit(1)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建HTTP客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		os.Exit(1)
	}

	// 打印响应状态码和内容
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应内容: %s\n", string(body))
}

func tags() {
	// 设置请求URL
	url := "http://localhost:11434/api/tags"

	// 创建HTTP客户端
	client := &http.Client{}

	// 创建HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 处理响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("请求失败，状态码: %d，响应: %s\n", resp.StatusCode, string(body))
		return
	}

	// 解析JSON响应
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return
	}

	// 格式化输出JSON
	formattedJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("格式化JSON失败:", err)
		return
	}

	// 输出结果
	fmt.Println(string(formattedJSON))
}

func main() {
	generate()
	tags()
}
