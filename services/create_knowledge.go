package srv

import (
	"context"

	"ksogit.kingsoft.net/o/xfx/validator"

	ctx "magic-box/context"
)

var _ = validator.Validate

func NewCreateKnowledgeService(moduleCtx *ctx.ModuleContext) CreateKnowledgeFunc {
	return func(ctx context.Context, request *CreateKnowledgeRequest) (*CreateKnowledgeResponse, error) {
		err := validator.Validate(request)
		if err != nil {
			return nil, err
		}
		out := NewCreateKnowledgeResponse()
		return &out, nil
	}
}
