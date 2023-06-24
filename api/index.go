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
  "<title>Live Stream on xxiao.org</title>" +
  "</head>" +
  "<body style=\"display: flex; align-items: center; flex-direction: column; font-size: 24px; font-family: 'Microsoft YaHei', sans-serif;\">"+
  "<h1><center>Huya, Douyu, Douyin, YY Live Agent</center></h1>"+
  "<p style=\"font-weight: bold;\"><center>Access path: <i>https://jkio.ml/platform/id</i></center></p>"+
  "<h6><center>Learn more about it from GitHub Repository: <i><a href=\"https://github.com/LeonoreShaw/iptv-go\">https://github.com/LeonoreShaw/iptv-go</a></i></center></h6>"+
  "<hr style=\"border: 1px; border-top: 2px solid blue;\">"+
  "<h1><center>虎牙、斗鱼、抖音、YY直播代理</center></h1>"+
  "<p style=\"font-weight: bold;\"><center>访问路径：<i>https://jkio.ml/platform/id</i></center></p>"+
  "<h6><center>在 GitHub 仓库了解更多信息：<i><a href=\"https://github.com/LeonoreShaw/iptv-go\">https://github.com/LeonoreShaw/iptv-go</a></i></center></h6>"+
  "</body>"+
  "</html>")
  return
}