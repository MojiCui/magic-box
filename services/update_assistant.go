package srv

import (
	"context"
	"time"

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
		out.Id = request.Id
		out.Name = request.Name
		out.Description = request.Description
		now := time.Now().Format(time.RFC3339)
		out.CreatedAt = now
		out.UpdatedAt = now
		return &out, nil
	}
}
