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

type TestGetSolutionTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestGetSolutionTestSuite) Test() {
	Convey("GetSolution", s.T(), func() {
		// service := NewGetSolutionService(s.ModuleContext())
	})
}

func TestGetSolution(t *testing.T) {
	testsuite.Run(t, new(TestGetSolutionTestSuite))
}
