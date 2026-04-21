package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewDeleteKnowledgeService(moduleCtx *ctx.ModuleContext) DeleteKnowledgeFunc {
	return func(ctx context.Context, request *DeleteKnowledgeRequest) error {
		err := validator.Validate(request)
		if err != nil {
			return err
		}
		return nil
	}
}
