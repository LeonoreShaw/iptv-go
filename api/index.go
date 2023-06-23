package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "
  
    <div style="text-align: center;">Huya, Douyu, Douyin, YY Live Agent</div>
    <div style="text-align: center;">Access path: https://jkio.ml/live/platform/id</div>
    <div style="text-align: center;">GitHub repository: https://github.com/leonoreshaw/iptv-go</div>
  
  ")
  return
}