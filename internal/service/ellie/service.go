package ellie

import (
	"context"
	"errors"
	stdl "log"
	"os"

	"github.com/sorohimm/misc/conf"
	"github.com/sorohimm/misc/grace"
	"github.com/sorohimm/misc/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

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

	l, err := log.NewZap(appConf.Log.Level, appConf.Log.EncType)
	if err != nil {
		stdl.Fatal(err)
	}
	logger := l.WithOptions(zap.AddStacktrace(zapcore.DPanicLevel)).
		Sugar().
		With("v", version, "built", built, "app", name)
	return log.CtxWithLogger(ctx, logger.Desugar())
}

func (o *Service) initAppConfig(ctx context.Context) context.Context {
	appConf := &config.Config{}
	if err := conf.New(appConf); err != nil {
		if errors.Is(err, conf.ErrHelp) {
			os.Exit(0)
		}
		stdl.Fatalf("failed to read app config: %v", err)
	}

	return config.WithContext(ctx, appConf)
}

func (o *Service) run(ctx context.Context) {
	logger := log.FromContext(ctx).Sugar()
	if err := o.RunGroup.Run(func(err error) {
		if err != nil {
			logger.Error(err)
		}
	}); err != nil {
		logger.Error("unexpected error", zap.Error(err))
	}
}
