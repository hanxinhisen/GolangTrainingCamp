package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := gin.Default()

	app.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"code": 0})
	})

	fakeError := make(chan struct{})
	app.GET("/fakeError", func(c *gin.Context) {
		// 模拟退出
		fakeError <- struct{}{}

	})

	g, ctx := errgroup.WithContext(context.Background())
	srv := &http.Server{
		Addr:    ":8899",
		Handler: app,
	}
	// 启动监听服务
	g.Go(func() error {
		return srv.ListenAndServe()
	})
	// 启动 退出服务
	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("出错退出")
		case <-fakeError:
			fmt.Println("模拟退出")
		}
		// 延时五秒等待,等待请求处理完,5秒后强制停止
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return srv.Shutdown(ctx)

	})
	// 启动监听信号量退出服务
	g.Go(func() error {
		exit := make(chan os.Signal)
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-exit:
			return errors.Errorf("主动退出程序")
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("%s", err)
	}

}
