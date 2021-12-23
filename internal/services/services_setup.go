package services

import (
	"github.com/isalikov/goapi/internal/env"
	"github.com/isalikov/goapi/internal/utils"
	goservice "github.com/isalikov/goservice/pkg/service"
)

func Setup(cfg *env.Config, logger *utils.Logger) *Clients {
	var err error
	var cc Clients

	if cc.GoService, err = goservice.Connect(cfg.GrpcGoService, 0); err != nil {
		logger.Fatal(err, err.Error())
	}

	return &cc
}
