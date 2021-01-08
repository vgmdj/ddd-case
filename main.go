package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vgmdj/ddd-case/infrastructure/config"
	"github.com/vgmdj/ddd-case/infrastructure/libs/group"
	"github.com/vgmdj/ddd-case/infrastructure/repository/persistence"
	"github.com/vgmdj/ddd-case/interfaces/httpfacade"
	"go.uber.org/zap"
)

func main() {
	cnf := config.GlobalEntity()

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic exit", zap.Any("err", err))
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		quitNow := false
		signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			s := <-c
			if quitNow {
				log.Println("automation receive twice quit signal and will exit right now ")
			}

			log.Println("automation get a quit signal and will wait all work finish then quit", zap.String("signal", s.String()))
			cancel()
			quitNow = true
		}
	}()
	g := group.WithCancel(ctx)

	// init http server
	httpServer := httpfacade.NewGinServer(&cnf.HTTPConfig)
	g.Go(httpServer.Launch)

	// init database
	persistence.InitGlobalRepo(&cnf.DBConfig)
	g.Go(persistence.GlobalRepo.Ping)

	log.Println("automation start")

	err := g.Wait()
	if err != nil {
		log.Println("program exit", zap.Error(err))
	}

}
