package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/go-kratos/kratos/v2"
	kratosconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	nacosconfig "github.com/go-kratos/nacos/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"gopkg.in/yaml.v3"

	"github.com/setcreed/kshop/app/user/service/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "kshop.user.service"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf          string
	RunEnv            string
	RegisterHost      string
	RegisterPort      uint64
	RegisterNamespace string
	configID          = "user"
)

type EnvData struct {
	ServiceName       string
	RunEnv            string
	RegisterHost      string
	RegisterPort      uint64
	RegisterNamespace string
}

func NewEnv(name, env, host, port, namespace string) *EnvData {
	serviceName := os.Getenv(name)
	runEnv := os.Getenv(env)
	registerHost := os.Getenv(host)
	registerPort := os.Getenv(port)
	registerNamespace := os.Getenv(namespace)

	report, _ := strconv.ParseInt(registerPort, 10, 64)

	return &EnvData{
		ServiceName:       serviceName,
		RunEnv:            runEnv,
		RegisterHost:      registerHost,
		RegisterPort:      uint64(report),
		RegisterNamespace: registerNamespace,
	}
}

func init() {
	env := NewEnv("SERVICE_NAME", "RUN_ENV", "REGISTER_HOST", "REGISTER_PORT", "REGISTER_NAMESPACE")
	if Name == "" {
		Name = env.ServiceName
	}
	RunEnv = env.RunEnv
	RegisterHost = env.RegisterHost
	RegisterPort = env.RegisterPort
	RegisterNamespace = env.RegisterNamespace
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	//conf
	var c kratosconfig.Config
	var source kratosconfig.Source

	if RunEnv == "" || RegisterHost == "" || RegisterPort == 0 || RegisterNamespace == "" {
		source = file.NewSource(flagconf)
	} else {
		client, e := getConfigClient(RegisterNamespace)
		if e != nil {
			panic(e)
		}
		source = nacosconfig.NewConfigSource(client, nacosconfig.Group("DEFAULT_GROUP"), nacosconfig.DataID(configID))
	}

	c = kratosconfig.New(
		kratosconfig.WithSource(
			source,
		),
		kratosconfig.WithDecoder(func(kv *kratosconfig.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(bc.Trace.Endpoint)))
	if err != nil {
		panic(err)
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("service", "user/service"),
		)),
	)

	app, cleanup, err := initApp(bc.Server, &rc, bc.Data, logger, tp)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func getConfigClient(namespaceID string) (config_client.IConfigClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(RegisterHost, RegisterPort),
	}

	cc := constant.ClientConfig{
		NamespaceId:         namespaceID, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// a more graceful way to create naming client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		return nil, err
	}
	return client, nil
}
