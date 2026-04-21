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
		out.Status = "healthy"
		out.Message = moduleCtx.Name() + " service is running"
		out.Version = "dc34006"
		return &out, nil
	}
}
