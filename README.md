# go_repeat

用 `chatgpt` 生成的一个简单文本去重工具，方便日常使用小工具

# Useage

go 编译出二进制

```go
go build -ldflags "-s -w" main.go
```

执行

```
Usage of go_repeat.exe:
  -f string
        Comma-separated list of input file names
  -o string
        Output file name
```

使用时，`-f` 可指定一个或者多个文件，例如：

```
./go_repeat.exe -f 1.txt,2.txt -o res.txt
```

