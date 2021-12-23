package services

import (
	proto "github.com/isalikov/goservice/pkg/api/grpc"
)

type Clients struct {
	GoService proto.GoServiceClient
}
