package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewGetAiService(moduleCtx *ctx.ModuleContext) GetAiFunc {
	return func(ctx context.Context, request *GetAiRequest) (*GetAiResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewGetAiResponse()
		return &out, nil
	}
}
