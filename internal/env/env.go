package env

type Config struct {
	Env						string	`env:"ENV,default=dev"`
	GrpcGoService 			string 	`env:"GRPC_GO_SERVICE,default=127.0.0.1:50000"`
	Port                	int		`env:"PORT,default=3001"`
	Release     			string  `env:"RELEASE,default=0.0.0"`
	SentryDsn				string	`env:"SENTRY_DSN,default=https://..."`
}
