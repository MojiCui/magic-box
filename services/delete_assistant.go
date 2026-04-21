package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewDeleteAssistantService(moduleCtx *ctx.ModuleContext) DeleteAssistantFunc {
	return func(ctx context.Context, request *DeleteAssistantRequest) error {
		err := validator.Validate(request)
		if err != nil {
			return err
		}
		return nil
	}
}
