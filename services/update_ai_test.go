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

type TestUpdateAiTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestUpdateAiTestSuite) Test() {
	Convey("UpdateAi", s.T(), func() {
		// service := NewUpdateAiService(s.ModuleContext())
	})
}

func TestUpdateAi(t *testing.T) {
	testsuite.Run(t, new(TestUpdateAiTestSuite))
}
