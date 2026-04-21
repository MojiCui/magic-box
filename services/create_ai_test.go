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

type TestCreateAiTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestCreateAiTestSuite) Test() {
	Convey("CreateAi", s.T(), func() {
		// service := NewCreateAiService(s.ModuleContext())
	})
}

func TestCreateAi(t *testing.T) {
	testsuite.Run(t, new(TestCreateAiTestSuite))
}
