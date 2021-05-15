package server

import (
	"api/api/cache"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
	"github.com/valyala/fasthttp"
)

// RequestHandler - represents a RequestHandler
type RequestHandler fasthttp.RequestHandler
type Service struct {
	URL                       string
	Name                      string
	MaxConnsPerHost           int
	MaxIdemponentCallAttempts int
	PostScoreURI              string
	GetLeaderBoardURI         string
	Client                    *fasthttp.Client
}

// NewClient -
func NewClient(option *Option) (*Service, error) {
	c := &Service{
		URL:  option.URL,
		Name: option.Name,
		Client: &fasthttp.Client{
			MaxConnsPerHost:           option.MaxConnsPerHost,
			MaxIdemponentCallAttempts: option.MaxIdemponentCallAttempts,
		},
		PostScoreURI:      option.PostScoreURI,
		GetLeaderBoardURI: option.GetLeaderBoardURI,
	}

	return c, nil
}

// ServiceName -
const ServiceName = "apitest"

var (
	redis *cache.DB
)

// Option - 設定檔
type Option struct {
	Address                   string
	Compress                  bool
	Config                    *Config
	URL                       string
	Name                      string
	MaxConnsPerHost           int
	MaxIdemponentCallAttempts int
	PostScoreURI              string
	GetLeaderBoardURI         string
}

// Start - 啟動伺服器
func Start(opt Option) {
	defer func() {
	}()

	initializeServer(opt.Config)
	TenMinutesTask()

	handler := processRequest
	if opt.Compress {
		handler = fasthttp.CompressHandler(handler)
	}
	server := &fasthttp.Server{
		Handler:                       handler,
		Name:                          ServiceName,
		TCPKeepalive:                  true,
		TCPKeepalivePeriod:            300 * time.Second,
		DisableKeepalive:              true,
		DisableHeaderNamesNormalizing: true,
	}
	log.Printf("%s listening on address %s\n", server.Name, opt.Address)
	if err := server.ListenAndServe(opt.Address); err != nil {
		log.Fatalf("Error in ListenAndServe: %s\n", err)
	}

	// Graceful shutdown -
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		// kill -SIGINT XXXX 或 Ctrl+c
		os.Interrupt,
		syscall.SIGINT, // register that too, it should be ok
		// os.Kill等同于syscall.Kill
		os.Kill,
		syscall.SIGKILL, // register that too, it should be ok
		// kill -SIGTERM XXXX
		syscall.SIGTERM,
	)
	s := <-ch
	log.Printf("s...%v\n", s)
	server.Shutdown()
	log.Printf("ServerClose...\n")
}

// initializeServer - 初始化伺服器
func initializeServer(config *Config) {
	var err error

	// Initialize Redis
	redis, err = cache.NewDB(&cache.DBOption{
		Addr:                  config.DB.Redis.Host,
		Password:              config.DB.Redis.Password,
		DB:                    config.DB.Redis.DB,
		PoolSize:              config.DB.Redis.PoolSize,
		RedisScriptDefinition: config.DB.Redis.RedisScriptDefinition,
		RedisScriptDB:         config.DB.Redis.RedisScriptDB,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Initialize Redis Success!\n")

	if err != nil {
		panic(err)
	}
}

func processRequest(ctx *fasthttp.RequestCtx) {
	path := ctx.Path()
	router := createDefaultRouter()

	if handler := router.GetRequestHandler(string(path)); handler != nil {
		handler(ctx)
	} else {
		ctx.NotFound()
	}
}

// TenMinutesTask - 每日排程
func TenMinutesTask() {

	mi := "0,10,20,30,40,50"
	ho := "0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23"

	c := cron.New()
	c.AddFunc("0 "+mi+" "+ho+" * * *", func() {
		redis.DelZsetAll(LeaderKey, "score")
	})
	c.Start()
	log.Printf("Initialize HourTask Success!\n")
}
