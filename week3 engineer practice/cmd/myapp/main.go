package main

import (
	"context"
	v1 "geek.xingx/egPractice/api/product/app/v1"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	service, err := InitArticleService()
	if err != nil {
		log.Panicf("init article service fail, %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterBlogServerServer(s, service)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		<-ctx.Done()
		log.Println("shutdown server")
		s.GracefulStop()
		return nil
	})

	g.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("receive os signal, %v", sig)
		}
	})

	g.Go(func() error {
		sock, err := net.Listen("tcp", ":8080")
		if err != nil {
			return errors.Wrapf(err, "start server on port 8080 failed")
		}
		log.Println("grpc server start listen on port 8080")
		return s.Serve(sock)
	})

	log.Printf("errgroup exiting, %v\n", g.Wait())
}
