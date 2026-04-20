package testsuitex

import ctx "magic-box/context"

type SuiteWithContext struct {
	moduleContext *ctx.ModuleContext
}

func (s *SuiteWithContext) ModuleContext() *ctx.ModuleContext {
	return s.moduleContext
}

func (s *SuiteWithContext) SetupSuite() {
	s.moduleContext = &ctx.ModuleContext{}
}
