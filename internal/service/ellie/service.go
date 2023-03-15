package ellie

import (
	"context"

	"github.com/sorohimm/misc/grace"
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
	return nil
}

func (o *Service) initAppConfig(ctx context.Context) context.Context {
	return nil
}

func (o *Service) run(ctx context.Context) {

}
