package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewCreateAssistantService(moduleCtx *ctx.ModuleContext) CreateAssistantFunc {
	return func(ctx context.Context, request *CreateAssistantRequest) (*CreateAssistantResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewCreateAssistantResponse()
		return &out, nil
	}
}
