package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"

	pb "prometheus-manager/api/strategy/v1/push"
)

type (
	IPushLogic interface {
		Strategies(ctx context.Context, req *pb.StrategiesRequest) (*pb.StrategiesReply, error)
		DeleteStrategies(ctx context.Context, req *pb.DeleteStrategiesRequest) (*pb.DeleteStrategiesReply, error)
	}

	PushService struct {
		pb.UnimplementedPushServer

		logger *log.Helper
		logic  IPushLogic
	}
)

var _ pb.PushServer = (*PushService)(nil)

func NewPushService(logic IPushLogic, logger log.Logger) *PushService {
	return &PushService{logic: logic, logger: log.NewHelper(log.With(logger, "module", loadModuleName))}
}

func (l *PushService) Strategies(ctx context.Context, req *pb.StrategiesRequest) (*pb.StrategiesReply, error) {
	ctx, span := otel.Tracer(pushModuleName).Start(ctx, "PushService.Strategies")
	defer span.End()
	return l.logic.Strategies(ctx, req)
}

func (l *PushService) DeleteStrategies(ctx context.Context, req *pb.DeleteStrategiesRequest) (*pb.DeleteStrategiesReply, error) {
	ctx, span := otel.Tracer(pushModuleName).Start(ctx, "PushService.DeleteStrategies")
	defer span.End()
	return l.logic.DeleteStrategies(ctx, req)
}
