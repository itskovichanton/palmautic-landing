package app

import (
	"fmt"
	"github.com/itskovichanton/core/pkg/core"
	"github.com/itskovichanton/core/pkg/core/app"
	"net/http"
)

type PalmApp struct {
	app.IApp

	Config *core.Config
}

func (c *PalmApp) Run() error {
	filesDir := c.Config.GetOnBaseWorkDir("landingfiles")
	http.Handle("/", http.FileServer(http.Dir(filesDir)))
	return http.ListenAndServe(fmt.Sprintf(":%v", c.Config.Server.Port), nil)
}
