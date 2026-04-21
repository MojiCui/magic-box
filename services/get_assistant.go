package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewGetAssistantService(moduleCtx *ctx.ModuleContext) GetAssistantFunc {
	return func(ctx context.Context, request *GetAssistantRequest) (*GetAssistantResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewGetAssistantResponse()
		return &out, nil
	}
}
