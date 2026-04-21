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

type TestGetAiTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestGetAiTestSuite) Test() {
	Convey("GetAi", s.T(), func() {
		// service := NewGetAiService(s.ModuleContext())
	})
}

func TestGetAi(t *testing.T) {
	testsuite.Run(t, new(TestGetAiTestSuite))
}
