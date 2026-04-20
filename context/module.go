package ctx

import (
	"ksogit.kingsoft.net/o/xfx"
)

type keyType int

var key keyType

type ModuleContext struct {
	AppContext xfx.App
	*backingServices
}

func (c *ModuleContext) Name() string {
	return "magic-box"
}

func (c *ModuleContext) Init(app xfx.App) {
	c.AppContext = app
	c.backingServices = newBackingServices(c)
}

func (c *ModuleContext) Term(app xfx.App) {
}

func Get(app xfx.App) *ModuleContext {
	return app.Module(key).(*ModuleContext)
}

func init() {
	xfx.RegisterModule(key, &ModuleContext{})
}
