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

type TestCreateAssistantTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestCreateAssistantTestSuite) Test() {
	Convey("CreateAssistant", s.T(), func() {
		// service := NewCreateAssistantService(s.ModuleContext())
	})
}

func TestCreateAssistant(t *testing.T) {
	testsuite.Run(t, new(TestCreateAssistantTestSuite))
}
