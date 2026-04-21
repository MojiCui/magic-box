package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewCreateSolutionService(moduleCtx *ctx.ModuleContext) CreateSolutionFunc {
	return func(ctx context.Context, request *CreateSolutionRequest) (*CreateSolutionResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewCreateSolutionResponse()
		return &out, nil
	}
}
