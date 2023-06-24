package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<body style=\"display: flex; align-items: center; justify-content: center; height: 100vh;\">"+
  "<h1 style=\"text-align: center;\">Huya, Douyu, &#x4F60;&#x597D;&#xFF0C;&#x4E16;&#x754C;&#xFF01; Douyin, YY Live Agent</h1>"+
  "<p style=\"text-align: center;\">Access path: <i>https://jkio.ml/platform/id</i></p> "+
  "<h6 style=\"text-align: center;\">Learn more about it from GitHub Repository: <i><a href=\"https://github.com/LeonoreShaw/iptv-go\">https://github.com/LeonoreShaw/iptv-go</a></i></h6>"+
  "</body>")
  return
}