package server

import (
	v1 "github.com/sparklee/abc-api/api/abc/v1"
	v2 "github.com/sparklee/abc-api/api/helloworld/v1"
	"github.com/sparklee/abc-api/internal/conf"
	"github.com/sparklee/abc-api/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, wordService *service.WordService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterWordServiceHTTPServer(srv, wordService)
	v2.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
