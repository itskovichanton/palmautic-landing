package app

import (
	"github.com/itskovichanton/core/pkg/core"
	"github.com/itskovichanton/core/pkg/core/app"
	"github.com/itskovichanton/server/pkg/server/di"
	"go.uber.org/dig"
)

type DI struct {
	di.DI
}

func (c *DI) InitDI() {

	container := dig.New()
	c.DI.InitDI(container)

	container.Provide(c.NewApp)
}

func (c *DI) NewApp(config *core.Config) app.IApp {
	return &PalmApp{
		Config: config,
	}
}
