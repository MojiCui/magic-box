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

type TestHealthCheckTestSuite struct {
	testsuite.Suite
	testsuitex.SuiteWithContext
}

func (s *TestHealthCheckTestSuite) Test() {
	Convey("HealthCheck", s.T(), func() {
		// service := NewHealthCheckService(s.ModuleContext())
	})
}

func TestHealthCheck(t *testing.T) {
	testsuite.Run(t, new(TestHealthCheckTestSuite))
}
