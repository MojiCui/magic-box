package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewDeleteAiService(moduleCtx *ctx.ModuleContext) DeleteAiFunc {
	return func(ctx context.Context, request *DeleteAiRequest) error {
		err := validator.Validate(request)
		if err != nil {
			return err
		}
		return nil
	}
}
