package util

import "context"

func TemplIsHtmx(ctx context.Context) bool {
	return ctx.Value("isHtmx").(bool)
}
