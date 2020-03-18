package main

import (
	"context"
	"flag"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example.com/go-microservice/user/handler"
	"example.com/go-microservice/user/service"
)

func main() {
	var httpAddr = flag.String("http", ":7001", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "user",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	// var db *sql.DB
	// {
	// 	var err error

	// 	db, err = sql.Open("postgres", dbsource)
	// 	if err != nil {
	// 		level.Error(logger).Log("exit", err)
	// 		os.Exit(-1)
	// 	}

	// }

	flag.Parse()
	ctx := context.Background()
	srv := &service.UserServiceImpl{}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := service.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := handler.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
