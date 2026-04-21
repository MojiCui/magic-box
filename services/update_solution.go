package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewUpdateSolutionService(moduleCtx *ctx.ModuleContext) UpdateSolutionFunc {
	return func(ctx context.Context, request *UpdateSolutionRequest) (*UpdateSolutionResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewUpdateSolutionResponse()
		return &out, nil
	}
}
