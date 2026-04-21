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

type TestCreateKnowledgeTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestCreateKnowledgeTestSuite) Test() {
	Convey("CreateKnowledge", s.T(), func() {
		// service := NewCreateKnowledgeService(s.ModuleContext())
	})
}

func TestCreateKnowledge(t *testing.T) {
	testsuite.Run(t, new(TestCreateKnowledgeTestSuite))
}
