package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	// 文件
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" +
		// object_name
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		// 错误响应
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.Close()
	// 写入文件
	io.Copy(f, r.Body)
}
