package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewUpdateAiService(moduleCtx *ctx.ModuleContext) UpdateAiFunc {
	return func(ctx context.Context, request *UpdateAiRequest) (*UpdateAiResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewUpdateAiResponse()
		return &out, nil
	}
}
