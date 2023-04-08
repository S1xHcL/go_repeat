package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 解析命令行参数
	fileNames := flag.String("f", "", "Comma-separated list of input file names")
	outputFile := flag.String("o", "", "Output file name")
	flag.Parse()

	// 打开输出文件
	outFile, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outFile.Close()

	// 记录所有出现过的行
	lines := make(map[string]bool)

	// 遍历所有输入文件
	for _, fileName := range splitFileNames(*fileNames) {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Error opening input file '%s': %v\n", fileName, err)
			continue
		}
		defer file.Close()

		// 逐行读取文件内容并去重
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if !lines[line] {
				fmt.Fprintln(outFile, line)
				lines[line] = true
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading input file '%s': %v\n", fileName, err)
		}
	}
}

// 将逗号分隔的文件名字符串分割成切片
func splitFileNames(names string) []string {
	if names == "" {
		return nil
	}
	return strings.Split(names, ",")
}