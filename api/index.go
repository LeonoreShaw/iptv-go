package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<div><h1>Huya, Douyu, Douyin, YY Live Agent</h1></div> <p>Access path: https://jkio.ml/platform/id</p> <p><h4>Learn more about it from GitHub Repository: https://github.com/LeonoreShaw/iptv-go</h4></p>")
  return
}