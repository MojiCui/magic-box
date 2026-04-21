package srv

import (
	"context"
	"time"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
	"magic-box/schemas"
)

var _ = validator.Validate

func NewListAisService(moduleCtx *ctx.ModuleContext) ListAisFunc {
	return func(ctx context.Context, request *ListAisRequest) (*ListAisResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewListAisResponse()
		item1 := schemas.NewAiItem()
		item1.Id = "ai-1"
		item1.Name = "测试AI1"
		item1.Description = "这是一个测试AI"
		now := time.Now().Format(time.RFC3339)
		item1.CreatedAt = now
		item1.UpdatedAt = now
		out.Items = append(out.Items, item1)
		out.Total = len(out.Items)
		return &out, nil
	}
}
