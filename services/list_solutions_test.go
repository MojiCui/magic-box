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

type TestListSolutionsTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestListSolutionsTestSuite) Test() {
	Convey("ListSolutions", s.T(), func() {
		// service := NewListSolutionsService(s.ModuleContext())
	})
}

func TestListSolutions(t *testing.T) {
	testsuite.Run(t, new(TestListSolutionsTestSuite))
}
