package options

import (
	"context"
	"fmt"
)

type demoConfig struct {
	demoInt int64
	demoBool bool
	demoString string
}

type Option func(*demoConfig)

func DemoInt(in int64) Option {
	return func(dc *demoConfig) {
		dc.demoInt = in
	}
}

func DemoBool(in bool) Option {
	return func(dc *demoConfig) {
		dc.demoBool = in
	}
}

func DemoString(in string) Option {
	return func(dc *demoConfig) {
		dc.demoString = in
	}
}

func (dc *demoConfig) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(dc)
	}
}

type demoConfigKey struct {}

func WithContext(ctx context.Context, cfg demoConfig) context.Context {
	return context.WithValue(ctx, demoConfigKey{}, cfg)
}


type createFunc func() demoConfig

func FromContextOrCreate(ctx context.Context, create createFunc) demoConfig {
	vc, ok := ctx.Value(demoConfigKey{}).(demoConfig)
	if !ok {
		return create()
	}
	return vc
}

type Service struct {}

func NewService() *Service {
	return &Service{}
}

func (s *Service) DefaultViewConfigCreater() createFunc {
	return func() demoConfig {
		return demoConfig{
			// default val
			demoString: "hello",
		}
	}
}
func (s *Service) Demo(ctx context.Context) {
	cfg := FromContextOrCreate(ctx, s.DefaultViewConfigCreater())
	// print: {27 false hello}
	fmt.Println(cfg)
}