package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewDeleteSolutionService(moduleCtx *ctx.ModuleContext) DeleteSolutionFunc {
	return func(ctx context.Context, request *DeleteSolutionRequest) error {
		err := validator.Validate(request)
		if err != nil {
			return err
		}
		return nil
	}
}
