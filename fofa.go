package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type FofaEntry struct {
	URL   string `json:"url"`
	Port  int    `json:"port"`
	Title string `json:"title"`
	IP    string `json:"ip"`
}

func main() {
	// 打开输入文件
	inputFile, err := os.Open("fofahack.txt")
	if err != nil {
		fmt.Println("无法打开输入文件:", err)
		return
	}
	defer inputFile.Close()

	// 打开输出文件
	outputFile, err := os.Create("ip.txt")
	if err != nil {
		fmt.Println("无法创建输出文件:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		
		// 将单引号替换为双引号
		line = strings.ReplaceAll(line, "'", "\"")

		var entry FofaEntry
		// 解析JSON行
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			fmt.Println("解析JSON失败:", err)
			continue
		}

		// 将IP和端口组合成期望的格式
		outputLine := fmt.Sprintf("%s:%d\n", entry.IP, entry.Port)
		// 写入输出文件
		if _, err := outputFile.WriteString(outputLine); err != nil {
			fmt.Println("写入输出文件失败:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取输入文件失败:", err)
	}
}
