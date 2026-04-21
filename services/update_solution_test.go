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

type TestUpdateSolutionTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestUpdateSolutionTestSuite) Test() {
	Convey("UpdateSolution", s.T(), func() {
		// service := NewUpdateSolutionService(s.ModuleContext())
	})
}

func TestUpdateSolution(t *testing.T) {
	testsuite.Run(t, new(TestUpdateSolutionTestSuite))
}
