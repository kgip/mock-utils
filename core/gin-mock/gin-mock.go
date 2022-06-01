package gin_mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
)

// DoRequest 初始化请求并且发送请求，返回响应结果
func DoRequest(router *gin.Engine, method, url string, data interface{}, headers ...map[string]string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var reader io.Reader
	if data != nil {
		body, _ := json.Marshal(data)
		reader = bytes.NewReader(body)
		if len(headers) > 0 {
			headers[0]["Content-Type"] = "application/json"
		} else {
			headers = append(headers, map[string]string{"Content-Type": "application/json"})
		}
	}
	req, _ := http.NewRequest(method, url, reader)
	for _, header := range headers {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	router.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	return w
}
