package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
	"magic-box/schemas"
)

var _ = validator.Validate

func NewListAssistantsService(moduleCtx *ctx.ModuleContext) ListAssistantsFunc {
	return func(ctx context.Context, request *ListAssistantsRequest) (*ListAssistantsResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewListAssistantsResponse()
		item1 := schemas.NewAssistantItem()
		item1.Id = "assistant-1"
		item1.Name = "测试助手1"
		item1.Description = "这是一个测试助手"
		item1.CreatedAt = "2024-01-01T00:00:00Z"
		item1.UpdatedAt = "2024-01-01T00:00:00Z"
		out.Items = append(out.Items, item1)
		out.Total = len(out.Items)
		return &out, nil
	}
}
