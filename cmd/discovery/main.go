package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/go-kratos/kratos/pkg/log"
	"github.com/peterlearn/discovery/conf"
	"github.com/peterlearn/discovery/discovery"
	"github.com/peterlearn/discovery/http"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.Log)
	dis, cancel := discovery.New(conf.Conf)
	http.Init(conf.Conf, dis)
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("discovery get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			cancel()
			time.Sleep(time.Second)
			log.Info("discovery quit !!!")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
