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

type TestDeleteAiTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestDeleteAiTestSuite) Test() {
	Convey("DeleteAi", s.T(), func() {
		// service := NewDeleteAiService(s.ModuleContext())
	})
}

func TestDeleteAi(t *testing.T) {
	testsuite.Run(t, new(TestDeleteAiTestSuite))
}
