package ellie

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/sorohimm/misc/grace"
	"github.com/sorohimm/misc/log"

	"github.com/sorohimm/ellie/internal/service/ellie/config"
)

func NewService() *Service {
	return &Service{
		RunGroup: grace.NewRunGroup(),
	}
}

type Service struct {
	*grace.RunGroup
}

func (o *Service) Init(name, version, built string) {
	var (
		ctx = context.Background()
	)

	// prepare ctx, configs and logger
	ctx = o.initAppConfig(ctx)
	ctx = o.initLogger(ctx, name, version, built)

	o.run(ctx)
}

func (o *Service) initLogger(ctx context.Context, name, version, built string) context.Context {
	appConf := config.FromContext(ctx)

	// init logger
	l, err := log.NewZap(appConf.Log.Level, appConf.Log.EncType)
	if err != nil {
		stdl.Fatal(err)
	}
	logger := l.WithOptions(
		zap.AddStacktrace(zapcore.DPanicLevel)).
		Sugar().
		With("v", version, "built", built, "app", name)
	return log.CtxWithLogger(ctx, logger.Desugar())
}

func (o *Service) initAppConfig(ctx context.Context) context.Context {

}

func (o *Service) run(ctx context.Context) {

}
