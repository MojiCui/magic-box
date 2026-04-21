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

type TestGetAssistantTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestGetAssistantTestSuite) Test() {
	Convey("GetAssistant", s.T(), func() {
		// service := NewGetAssistantService(s.ModuleContext())
	})
}

func TestGetAssistant(t *testing.T) {
	testsuite.Run(t, new(TestGetAssistantTestSuite))
}
