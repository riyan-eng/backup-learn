package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"gohertz/config"
	"gohertz/env"
	"gohertz/infrastructure"
	"gohertz/internal/repository"
	"gohertz/internal/router"
	"gohertz/middleware"
	"log"
	"runtime"
	"time"

	_ "gohertz/docs"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	http2Config "github.com/hertz-contrib/http2/config"
	"github.com/hertz-contrib/http2/factory"
	"github.com/hertz-contrib/swagger"     // hertz-swagger middleware
	swaggerFiles "github.com/swaggo/files" // swagger embed files
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU - 1)
	}
	env.LoadEnvironmentFile()
	env.NewEnv()

	config.NewLimiterStore()
	config.NewLogger()

	infrastructure.ConnectSqlDB()
	infrastructure.ConnectSqlxDB()
	infrastructure.ConnRedis()
	infrastructure.NewLocalizer()
}

// @title HertzTest
// @version 1.0
// @description This is a demo using Hertz.

// @contact.name hertz-contrib
// @contact.url https://github.com/hertz-contrib

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer access token here
func main() {
	// http2 config
	cfg := &tls.Config{
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}
	cert, err := tls.LoadX509KeyPair(env.NewEnv().CERTIFICATE_CRT, env.NewEnv().CERTIFICATE_KEY)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Certificates = append(cfg.Certificates, cert)

	// create instance
	app := server.Default(
		server.WithHostPorts(env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT),
		server.WithALPN(true),
		server.WithTLS(cfg),
	)
	defer app.Spin()

	// http2 protocol
	app.AddProtocol("h2", factory.NewServerFactory(
		http2Config.WithReadTimeout(time.Minute),
		http2Config.WithDisableKeepAlive(false),
	))
	cfg.NextProtos = append(cfg.NextProtos, "h2")

	app.GET("/ping", PingHandler)

	// swagger
	// url := swagger.DocExpansion() // The url pointing to API definition
	app.GET("/docs/*any", swagger.WrapHandler(swaggerFiles.Handler))

	// middleware
	app.Use(recovery.Recovery())
	app.Use(middleware.RequestId())
	app.Use(middleware.Logger())
	app.Use(middleware.Limiter())
	app.Use(infrastructure.LocalizerMiddleware())

	// service
	dao := repository.NewDAO(infrastructure.SqlDB, infrastructure.SqlxDB, infrastructure.Redis, config.NewEnforcer())

	// router
	routers := router.NewRouter(app, &dao)
	routers.Index()
	routers.Authentication()
	routers.Example()

	// startup log
	fmt.Println("server run on:", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
}

func PingHandler(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(200, map[string]string{
		"ping": "pong",
	})
}
