package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	statsd "github.com/asppj/thanos-compose/statsd-demo/statsd"
)

var (
	addr   = "127.0.0.1:19125"
	prefix = "statsd_demo_liusp"
)

func main() {
	_addr := os.Getenv("STATSD_ADDR")
	_prefix := os.Getenv("STATSD_PREFIX")
	if _addr != "" {
		addr = _addr
	}
	if _prefix != "" {
		prefix = _prefix
	}
	rand.Seed(time.Now().UnixNano())
	statsd.InitStatsdClient(prefix, addr)
	defer statsd.CloseStatsdClient()
	// 超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	// 取消
	watch(func() {
		cancel()
	})
	go NewRandRunner().LooupDo(ctx, func() {
	})
	// statsd
	fns := []func(){
		func() {
			statsd.ReportSuccessResponse(randIndex(200, 201, 301).(int), randIndex("/path", "/index.html", "/home").(string))
		},
		func() {
			statsd.ReportCostTimeResponse(randIndex("/path", "/index.html", "/home").(string), randIndex("GET", "POST", "PUT").(string), time.Second*time.Duration(randIndex(1, 2, 3, 4, 5).(int)))
		},
		func() {
			statsd.ReportCashNum(randIndex(3.4, 1.444, 0.434, 5.0, 33.0).(float64))
		},
		func() {
			statsd.ReportErrorResponse(randIndex(404, 502, 501).(int), randIndex("/path", "/index.html", "/home").(string))
		},
		func() {
			statsd.ReportIgnoreErrorResponse(randIndex(404, 502, 501).(int), randIndex("/path", "/index.html", "/home").(string))
		},
		func() {
			statsd.ReportFeastResponse(randIndex(200, 201, 301, 500, 501, 404).(int), randIndex("main", "ranking").(string))
		},
		func() {
			statsd.ReportRedisError(randIndex("main", "ranking").(string))
		},
		func() {
			statsd.ReportOutResponse(randIndex(200, 201, 301, 500, 501, 404).(int), randIndex("main", "ranking").(string))
		},
		statsd.ReportNormalSuccessResponse,
		statsd.ReportNormalErrorResponse,
		statsd.ReportWithdrawTaskletSuccessPassResponse,
	}
	loop(ctx, fns...)

	<-ctx.Done()
}

func watch(fn func()) {
	sc := make(chan os.Signal, 3)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR2)
	go func() {
		for {
			<-sc
			fn()
		}
	}()
}

type randRunner struct {
	min int // 100 Millisecond
	max int // 100 Millisecond
}

func NewRandRunner() *randRunner {
	return &randRunner{min: 1, max: 10}
}

func (r *randRunner) LooupDo(ctx context.Context, fn func()) {
	tick := time.NewTicker(time.Duration(r.min) * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			r.roundSleep()
			fn()
			fmt.Printf("%#v\n", fn)
		}
	}
}

func (r randRunner) roundSleep() {
	n := rand.Int31n(int32(r.max - r.min))
	time.Sleep(time.Duration(n) * 1000 * time.Millisecond)
}

func randIndex(args ...interface{}) interface{} {
	n := len(args)
	if n == 0 {
		return 0
	}
	r := rand.Int31n(int32(n))
	return args[r]
}

func loop(ctx context.Context, fns ...func()) {
	for _, fn := range fns {
		go NewRandRunner().LooupDo(ctx, fn)
	}
}
