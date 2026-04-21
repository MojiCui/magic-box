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

type TestUpdateAssistantTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestUpdateAssistantTestSuite) Test() {
	Convey("UpdateAssistant", s.T(), func() {
		// service := NewUpdateAssistantService(s.ModuleContext())
	})
}

func TestUpdateAssistant(t *testing.T) {
	testsuite.Run(t, new(TestUpdateAssistantTestSuite))
}
