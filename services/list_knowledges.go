package srv

import (
	"context"
	"time"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
	"magic-box/schemas"
)

var _ = validator.Validate

func NewListKnowledgesService(moduleCtx *ctx.ModuleContext) ListKnowledgesFunc {
	return func(ctx context.Context, request *ListKnowledgesRequest) (*ListKnowledgesResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewListKnowledgesResponse()
		item1 := schemas.NewKnowledgeItem()
		item1.Id = "knowledge-1"
		item1.Name = "测试知识1"
		item1.Description = "这是一个测试知识"
		now := time.Now().Format(time.RFC3339)
		item1.CreatedAt = now
		item1.UpdatedAt = now
		out.Items = append(out.Items, item1)
		out.Total = len(out.Items)
		return &out, nil
	}
}
