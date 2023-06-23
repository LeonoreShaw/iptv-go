package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "
  
    <div>Huya, Douyu, Douyin, YY Live Agent</div>
    <div>Access path: https://jkio.ml/live/platform/id</div>
    <div>GitHub repository: https://github.com/leonoreshaw/iptv-go</div>
  
  ")
  return
}