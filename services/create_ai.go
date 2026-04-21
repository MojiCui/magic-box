package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewCreateAiService(moduleCtx *ctx.ModuleContext) CreateAiFunc {
	return func(ctx context.Context, request *CreateAiRequest) (*CreateAiResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewCreateAiResponse()
		return &out, nil
	}
}
