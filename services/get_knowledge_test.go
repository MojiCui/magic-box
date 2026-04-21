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

type TestGetKnowledgeTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestGetKnowledgeTestSuite) Test() {
	Convey("GetKnowledge", s.T(), func() {
		// service := NewGetKnowledgeService(s.ModuleContext())
	})
}

func TestGetKnowledge(t *testing.T) {
	testsuite.Run(t, new(TestGetKnowledgeTestSuite))
}
