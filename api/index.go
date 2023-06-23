package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "
  
    <h2>Huya, Douyu, Douyin, YY Live Agent</h2>
    <h4>Access path: https://jkio.ml/live/platform/id</h4>
    <h4>GitHub repository: https://github.com/leonoreshaw/iptv-go</h4>
  
  ")
  return
}