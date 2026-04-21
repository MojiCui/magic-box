package srv

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"ksogit.kingsoft.net/o/xfx/monkey"
	"ksogit.kingsoft.net/o/xfx/testsuite"
	ctx "magic-box/context"
	"magic-box/testsuitex"
)

var _ = context.Canceled
var _ = monkey.Patches{}
var _ = ctx.ModuleContext{}

type TestUpdateKnowledgeTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestUpdateKnowledgeTestSuite) Test() {
	Convey("UpdateKnowledge", s.T(), func() {
		// service := NewUpdateKnowledgeService(s.ModuleContext())
	})
}

func TestUpdateKnowledge(t *testing.T) {
	testsuite.Run(t, new(TestUpdateKnowledgeTestSuite))
}
