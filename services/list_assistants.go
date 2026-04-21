package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewListAssistantsService(moduleCtx *ctx.ModuleContext) ListAssistantsFunc {
	return func(ctx context.Context, request *ListAssistantsRequest) (*ListAssistantsResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewListAssistantsResponse()
		return &out, nil
	}
}
