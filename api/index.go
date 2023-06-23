package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1 text-align:center>虎牙、斗鱼、、抖音、YY直播代理</h1><center>访问路径：https://jkio.ml/live/<平台>/id</center><h2 text-align:center>GitHub仓库地址：https://github.com/leonoreshaw/iptv-go/</center>")
  return
}