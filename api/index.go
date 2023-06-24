package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<body style=\"display: flex; align-items: center; flex-direction: column;\"><h1 style=\"text-align: center;\">Huya, Douyu, Douyin, YY Live Agent</h1> <p style=\"text-align: center;\"><i>Access path: https://jkio.ml/platform/id</i></p> <h6 style=\"text-align: center;\"><i>Learn more about it from GitHub Repository: https://github.com/LeonoreShaw/iptv-go</i></h6></body>")
  return
}