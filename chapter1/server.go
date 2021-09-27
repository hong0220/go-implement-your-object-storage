package main

import (
	"./objects"
	"log"
	"net/http"
	"os"
)

/**
 * os.File 实现 io.Writer 和 io.Reader

 * 启动：
 * LISTEN_ADDRESS=:12345 STORAGE_ROOT=/tmp go run server.go
 *
 * 创建：
 * mkdir /tmp/object
 *
 * 测试：
 * curl 127.0.0.1:12345/objects/test -XPUT -d"this is a test object"
 * curl 127.0.0.1:12345/objects/test
 */
func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
