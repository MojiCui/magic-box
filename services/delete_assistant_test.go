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

type TestDeleteAssistantTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestDeleteAssistantTestSuite) Test() {
	Convey("DeleteAssistant", s.T(), func() {
		// service := NewDeleteAssistantService(s.ModuleContext())
	})
}

func TestDeleteAssistant(t *testing.T) {
	testsuite.Run(t, new(TestDeleteAssistantTestSuite))
}
