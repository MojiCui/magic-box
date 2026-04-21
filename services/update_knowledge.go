package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewUpdateKnowledgeService(moduleCtx *ctx.ModuleContext) UpdateKnowledgeFunc {
	return func(ctx context.Context, request *UpdateKnowledgeRequest) (*UpdateKnowledgeResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewUpdateKnowledgeResponse()
		return &out, nil
	}
}
