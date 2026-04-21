package srv

import (
	"context"
	"fmt"
	"time"

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
		out.Name = request.Name
		out.Description = request.Description
		out.Id = fmt.Sprintf("assistant-%d", time.Now().Unix())
		now := time.Now().Format(time.RFC3339)
		out.CreatedAt = now
		out.UpdatedAt = now
		return &out, nil
	}
}
