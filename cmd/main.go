package main

import (
	"context"
	"params-with-ctx/options"
)

func main() {
	svc := options.NewService()
	cfg := svc.DefaultViewConfigCreater()()
	opts := []options.Option{
		options.DemoInt(27),
	}
	cfg.Apply(opts...)
	ctx := context.Background()
	ctx = options.WithContext(ctx, cfg)

	svc.Demo(ctx)
}
