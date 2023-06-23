package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h2>Huya, Douyu, Douyin, YY Live Agent</h2><h4>Access path</h4><h4>GitHub repository</h4>
  
  ")
  return
}