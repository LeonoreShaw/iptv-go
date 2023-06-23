package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<center><h1>Huya, Douyu, Douyin, YY Live Agent</h1></center> <center><p>Access path: <i>https://jkio.ml/platform/id</i></p></center> <center><p><h6>Learn more about it from GitHub Repository: <i>https://github.com/LeonoreShaw/iptv-go</i></h6></p></center>")
  return
}