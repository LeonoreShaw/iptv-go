package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1 align="center">Huya, Douyu, , Douyin, YY Live Agent</h1><center>Access path: https://jkio.ml/live/<platform>/id</center><h2 align="center">GitHub repository: https://github.com/leonoreshaw/iptv-go/</center>")
  return
}