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

type TestListAisTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestListAisTestSuite) Test() {
	Convey("ListAis", s.T(), func() {
		// service := NewListAisService(s.ModuleContext())
	})
}

func TestListAis(t *testing.T) {
	testsuite.Run(t, new(TestListAisTestSuite))
}
