package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewGetKnowledgeService(moduleCtx *ctx.ModuleContext) GetKnowledgeFunc {
	return func(ctx context.Context, request *GetKnowledgeRequest) (*GetKnowledgeResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewGetKnowledgeResponse()
		return &out, nil
	}
}
