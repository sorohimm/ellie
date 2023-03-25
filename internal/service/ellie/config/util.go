package config

import (
	"context"
	"strings"
)

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

func splitStringBy(str, div, cutset string) []string {
	ret := make([]string, 0)
	for _, s := range strings.Split(str, div) {
		cs := strings.Trim(s, cutset)
		if len(cs) > 0 {
			ret = append(ret, cs)
		}
	}
	return ret
}
