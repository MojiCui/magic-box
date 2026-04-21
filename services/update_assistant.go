package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewUpdateAssistantService(moduleCtx *ctx.ModuleContext) UpdateAssistantFunc {
	return func(ctx context.Context, request *UpdateAssistantRequest) (*UpdateAssistantResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewUpdateAssistantResponse()
		return &out, nil
	}
}
