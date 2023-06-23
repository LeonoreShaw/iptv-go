package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<div><h1>Huya, Douyu, Douyin, YY Live Agent</h1></div><p><h4>Access path</h4></p><p><h4>GitHub repository</h4></p>")
  return
}