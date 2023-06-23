package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1  style="text-align: center">Huya, Douyu, , Douyin, YY Live Agent</h1><center>Access path: https://jkio.ml/live/<platform>/id</center> <center">GitHub repository: <a href="https://github.com/leonoreshaw/iptv-go/"> https://github.com/leonoreshaw/iptv-go/ </a></center>")
  return
}