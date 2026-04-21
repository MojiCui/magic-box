package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewHealthCheckService(moduleCtx *ctx.ModuleContext) HealthCheckFunc {
	return func(ctx context.Context, request *HealthCheckRequest) (*HealthCheckResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewHealthCheckResponse()
		return &out, nil
	}
}
