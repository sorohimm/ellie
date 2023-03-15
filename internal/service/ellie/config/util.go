package config

import "context"

type confKey struct{} // or exported to use outside the package

func WithContext(ctx context.Context, c *Config) context.Context {
	return context.WithValue(ctx, confKey{}, c)
}

func FromContext(ctx context.Context) *Config {
	if cc, ok := ctx.Value(confKey{}).(*Config); ok {
		return cc
	}
	return NewDefaultConfig()
}

func NewDefaultConfig() *Config {
	return &Config{}
}
