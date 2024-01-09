package data

import (
	"context"

	"prometheus-manager/app/prom_agent/internal/biz/bo"
	"prometheus-manager/app/prom_agent/internal/biz/repository"

	"github.com/go-kratos/kratos/v2/log"
)

type PingRepo struct {
	data *Data
	log  *log.Helper
}

func (l *PingRepo) Ping(ctx context.Context, g *bo.Ping) (*bo.Ping, error) {
	return &bo.Ping{Hello: "world"}, nil
}

// NewPingRepo .
func NewPingRepo(data *Data, logger log.Logger) repository.PingRepo {
	return &PingRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data.Ping")),
	}
}
