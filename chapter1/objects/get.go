package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {
	// 文件
	f, e := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" +
		// object_name
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		// 错误响应
		w.WriteHeader(http.StatusNotFound)
		return
	}

	defer f.Close()
	io.Copy(w, f)
}
