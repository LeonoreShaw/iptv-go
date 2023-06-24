package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<!DOCTYPE html>"+
  "<html>"+
  "<head>" +
  "<meta charset=\"UTF-8\">" +
  "<title>My Page</title>" +
  "</head>" +
  "<body style=\"display: flex; align-items: center; flex-direction: column; font-size: 24px; font-family: 'Microsoft YaHei', sans-serif;\">"+
  "<h1><center>Huya, Douyu, Douyin, YY Live Agent</center></h1>"+
  "<p style=\"font-weight: bold;\"><center>Access path: <i>https://jkio.ml/platform/id</i></center></p>"+
  "<h6><center>Learn more about it from GitHub Repository: <i><a href=\"https://github.com/LeonoreShaw/iptv-go\">https://github.com/LeonoreShaw/iptv-go</a></i></center></h6>"+
  "<hr>"+
  "<h1><center>Huya, Douyu, &#x4F60;&#x597D;&#xFF0C;&#x4E16;&#x754C;&#xFF01; Douyin, YY Live Agent</center></h1>"+
  "<p style=\"font-weight: bold;\"><center>Access path: <i>https://jkio.ml/platform/id</i></center></p>"+
  "<h6><center>中文Learn more about it from GitHub Repository: <i><a href=\"https://github.com/LeonoreShaw/iptv-go\">https://github.com/LeonoreShaw/iptv-go</a></i></center></h6>"+
  "</body>"+
  "</html>")
  return
}