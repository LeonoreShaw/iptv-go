package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<body style=\"display: flex; justify-content: center \">"+
  "<h1>Huya, Douyu, &#x4F60;&#x597D;&#xFF0C;&#x4E16;&#x754C;&#xFF01; Douyin, YY Live Agent</h1>"+
  "<p style=\"font-weight: bold;\">Access path: <i>https://jkio.ml/platform/id</i></p>"+
  "<h6>Learn more about it from GitHub Repository: <i><a href=\"https://github.com/LeonoreShaw/iptv-go\">https://github.com/LeonoreShaw/iptv-go</a></i></h6>"+
  "</body>")
  return
}