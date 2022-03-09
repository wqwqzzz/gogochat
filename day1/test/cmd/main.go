package main
import (
    "gochat/internal/route"
    "time"
    "net/http"
    
)
func main(){
    
    newRouter := route.NewRoute()
    s := &http.Server{
    	Addr:           ":8888",
    	Handler:        newRouter,
    	ReadTimeout:    10 * time.Second,
    	WriteTimeout:   10 * time.Second,
    	MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
    
    
}