package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<div><h4>Huya, Douyu, Douyin, YY Live Agent</h4></div><p>Access path: <a href:"https://baidu.com">https://jkio.ml/<platform>/id</a></p><p><h4>GitHub repository</h4></p>")
  return
}