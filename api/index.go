package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "
  <div class="container">
    <div style="text-align: center;">Huya, Douyu, Douyin, YY Live Agent</div>
    <div style="text-align: center;">Access path: <a href="https://jkio.ml/live/platform/id">https://jkio.ml/live/platform/id</a></div>
    <div style="text-align: center;">GitHub repository: <a href="https://github.com/leonoreshaw/iptv-go">https://github.com/leonoreshaw/iptv-go</a></div>
  </div>
  ")
  return
}