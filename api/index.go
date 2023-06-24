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
  "<body style=\"font-size: 24px; \">"+
  "<h1><center>Huya, Douyu, Douyin, YY Live Agent</center></h1>"+
  "<p style=\"font-weight: bold;\"><center>Access path: <i>https://jkio.ml/platform/id</i></center></p>"+
  "<h6><center>Learn more about it from GitHub Repository: <i><a href=\"https://github.com/LeonoreShaw/iptv-go\" target=\"_blank\">https://github.com/LeonoreShaw/iptv-go</a></i></center></h6>"+
  "<hr style=\"border: 5px; border-top: 2px solid blue;\">"+
  "<h1><center>虎牙、斗鱼、抖音、YY直播代理</center></h1>"+
  "<p style=\"font-weight: bold;\"><center>访问路径：<i>https://jkio.ml/platform/id</i></center></p>"+
  "<h6><center>在 GitHub 仓库了解更多信息：<i><a href=\"https://github.com/LeonoreShaw/iptv-go\" target=\"_blank\">https://github.com/LeonoreShaw/iptv-go</a></i></center></h6>"+
  "<p><center>&copy; 2020-2023 <a href=\"https://xxiao.org\" target=\"_blank\">Léonore</a> All rights reserved.</center></p>"+
  "</body>"+
  "</html>")
  return
}