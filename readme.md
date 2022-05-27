# 基于 go 实现 wasm 解析 apk 的思路尝试

## go 语言本地实现

> 通过使用 `─ go run main.go`的方式 ，打印结果如下

```bash
# go run main.go                                                                                                                                                         
2022/05/27 11:49:27   icon %!q(<nil>)...
2022/05/27 11:49:27   pkgName "com.microsoft.windowsintune.companyportal"...
2022/05/27 11:49:27   appLabel "Company Portal"...

```



> 通过`GOOS=js GOARCH=wasm go build -o libs.wasm main.go `打包成wasm应用之后 

在web运行的效果 直接报错 

```bash
11:51:21.598 wasm_exec.js:22 panic: runtime error: invalid memory address or nil pointer dereference
11:51:21.599 wasm_exec.js:22 	panic: runtime error: invalid memory address or nil pointer dereference
11:51:21.599 wasm_exec.js:22 [signal 0xb code=0x0 addr=0x0 pc=0x0]
11:51:21.599 wasm_exec.js:22 
11:51:21.599 wasm_exec.js:22 goroutine 1 [running]:
11:51:21.600 wasm_exec.js:22 github.com/shogo82148/androidbinary/apk.(*Apk).Close(0x0)
11:51:21.600 wasm_exec.js:22 	/Users/apanda/MEGAsync/Projects/GoProjects/pkg/mod/github.com/shogo82148/androidbinary@v1.0.3/apk/apk.go:69 +0x2
11:51:21.600 wasm_exec.js:22 panic({0x1b740, 0x174ab0})
11:51:21.600 wasm_exec.js:22 	/usr/local/Cellar/go/1.18.2/libexec/src/runtime/panic.go:838 +0x29
11:51:21.600 wasm_exec.js:22 github.com/shogo82148/androidbinary/apk.(*Apk).Icon(0x0, 0x0)
11:51:21.601 wasm_exec.js:22 	/Users/apanda/MEGAsync/Projects/GoProjects/pkg/mod/github.com/shogo82148/androidbinary@v1.0.3/apk/apk.go:77 +0x2
11:51:21.601 wasm_exec.js:22 main.main()
11:51:21.601 wasm_exec.js:22 	/Users/apanda/MEGAsync/Projects/GoProjects/go_wasm/main.go:13 +0x3
```

## 结论 此路不通

# 参考文档

> [文档](https://segmentfault.com/a/1190000024501442)

官方文档 ：https://github.com/golang/go/wiki/WebAssembly#getting-started
