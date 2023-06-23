package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1><center>Huya, Douyu, Douyin, YY Live Agent</center></h1> <p><center>Access path: <i>https://jkio.ml/platform/id</i></center></p> <h6><center>Learn more about it from GitHub Repository: <i>https://github.com/LeonoreShaw/iptv-go</i></center></h6>")
  return
}