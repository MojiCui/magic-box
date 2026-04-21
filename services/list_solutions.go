package srv

import (
	"context"
	"time"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
	"magic-box/schemas"
)

var _ = validator.Validate

func NewListSolutionsService(moduleCtx *ctx.ModuleContext) ListSolutionsFunc {
	return func(ctx context.Context, request *ListSolutionsRequest) (*ListSolutionsResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewListSolutionsResponse()
		item1 := schemas.NewSolutionItem()
		item1.Id = "solution-1"
		item1.Name = "测试方案1"
		item1.Description = "这是一个测试方案"
		now := time.Now().Format(time.RFC3339)
		item1.CreatedAt = now
		item1.UpdatedAt = now
		out.Items = append(out.Items, item1)
		out.Total = len(out.Items)
		return &out, nil
	}
}
