package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/IRFAN374/upsvc/repository/token"
	// "github.com/IRFAN374/upsvc/repository/token"


	"github.com/go-kit/kit/log"
	gokitlogzap "github.com/go-kit/kit/log/zap"
	"github.com/gorilla/mux"
	"github.com/oklog/oklog/pkg/group"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var env string

func main() {

	fmt.Println("............user profile Service.............")

	flag.Parse()

	if env == "" {
		os.Getenv("env")
	}

	// serviceName := fmt.Sprintf("%s-user_profile_service", env)

	// logger

	// debugLogger, _, _, _ := getLogger(serviceName, zapcore.DebugLevel)

	// HTTP middleware
	var mwf []mux.MiddlewareFunc

	// router
	httpRouter := mux.NewRouter().StrictSlash(false)
	httpRouter.Use(mwf...)

	// Health Check Api
	httpRouter.PathPrefix("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})


	// token repository
	// var tokenRepo token.Repository
	// {
	// 	tokenRepo = 
	// }

	// avatar repository

	// avatar service

	// user repository

	// user service

	

	var server group.Group
	{
		httpServer := &http.Server{
			Addr: ":9000",
			// Handler: httpRouter,
		}

		server.Add(func() error {
			fmt.Printf("starting http Server on port : %d \n", 9000)
			return httpServer.ListenAndServe()
		}, func(err error) {
			// write code here for graceful shutdown

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
			defer cancel()

			httpServer.Shutdown(ctx)
		})

	}
	// interuption handling
	{

		cancelInterrupt := make(chan struct{})

		server.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(err error) {
			close(cancelInterrupt)
		})

	}

	fmt.Printf("exiting... errors: %v \n", server.Run())

}

func getLogger(serviceName string, level zapcore.Level) (debugL, infoL, errorL log.Logger, zapLogger *zap.Logger) {

	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	ec.EncodeDuration = zapcore.StringDurationEncoder
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(ec)

	fw, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}

	core := zapcore.NewCore(encoder, fw, level)
	zapLogger = zap.New(core)

	debugL = gokitlogzap.NewZapSugarLogger(zapLogger, zap.DebugLevel)
	debugL = log.With(debugL, "service", serviceName)

	infoL = gokitlogzap.NewZapSugarLogger(zapLogger, zap.InfoLevel)
	infoL = log.With(infoL, "service", serviceName)

	errorL = gokitlogzap.NewZapSugarLogger(zapLogger, zap.ErrorLevel)
	errorL = log.With(errorL, "service", serviceName)

	return
}
