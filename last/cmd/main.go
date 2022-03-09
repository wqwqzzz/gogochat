package main
import (
    "gochat/internal/route"
    "gochat/internal/server"
    "time"
    "net/http"
    "gochat/config"
    "gochat/pkg/global/log"
///	"gochat/internal/kafka"    
)
func main(){
    log.Logger.Info("config", log.Any("config", config.GetConfig()))
    newRouter := route.NewRoute()
    log.Logger.Info("start server", log.String("start", "start web sever in port 10086 ..."))
    go server.MyServer.Start()
    s := &http.Server{

    	Addr:           ":10086",
    	Handler:        newRouter,
    	ReadTimeout:    10 * time.Second,
    	WriteTimeout:   10 * time.Second,
    	MaxHeaderBytes: 1 << 20,
    }
    
    err :=s.ListenAndServe()
    if nil != err {
		log.Logger.Error("server error", log.Any("serverError", err))
	}
    
    
}
