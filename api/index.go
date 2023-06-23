package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Huya, Douyu, Douyin, YY Live Agent</h1> <p>Access path: <i>https://jkio.ml/platform/id</i></p> <p><h4>Learn more about it from GitHub Repository: <i>https://github.com/LeonoreShaw/iptv-go</i></h4></p>")
  return
}