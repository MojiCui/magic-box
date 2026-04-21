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

type TestCreateSolutionTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestCreateSolutionTestSuite) Test() {
	Convey("CreateSolution", s.T(), func() {
		// service := NewCreateSolutionService(s.ModuleContext())
	})
}

func TestCreateSolution(t *testing.T) {
	testsuite.Run(t, new(TestCreateSolutionTestSuite))
}
