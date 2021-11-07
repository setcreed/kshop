package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	v1 "github.com/setcreed/kshop/api/shop/interface/v1"
	"github.com/setcreed/kshop/app/shop/interface/internal/conf"
	"github.com/setcreed/kshop/app/shop/interface/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, s *service.ShopInterface, logger log.Logger, tp *tracesdk.TracerProvider) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
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
	v1.RegisterShopInterfaceHTTPServer(srv, s)
	return srv
}
