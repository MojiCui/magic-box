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

type TestListKnowledgesTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestListKnowledgesTestSuite) Test() {
	Convey("ListKnowledges", s.T(), func() {
		// service := NewListKnowledgesService(s.ModuleContext())
	})
}

func TestListKnowledges(t *testing.T) {
	testsuite.Run(t, new(TestListKnowledgesTestSuite))
}
