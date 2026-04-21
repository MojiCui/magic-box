package srv

import (
	"context"
	"time"

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
		out.Id = request.AssistantId
		out.Name = "测试助手"
		out.Description = "这是一个测试助手"
		now := time.Now().Format(time.RFC3339)
		out.CreatedAt = now
		out.UpdatedAt = now
		return &out, nil
	}
}
