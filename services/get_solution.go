package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewGetSolutionService(moduleCtx *ctx.ModuleContext) GetSolutionFunc {
	return func(ctx context.Context, request *GetSolutionRequest) (*GetSolutionResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewGetSolutionResponse()
		return &out, nil
	}
}
