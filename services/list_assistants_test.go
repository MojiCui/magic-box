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

type TestListAssistantsTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestListAssistantsTestSuite) Test() {
	Convey("ListAssistants", s.T(), func() {
		// service := NewListAssistantsService(s.ModuleContext())
	})
}

func TestListAssistants(t *testing.T) {
	testsuite.Run(t, new(TestListAssistantsTestSuite))
}
