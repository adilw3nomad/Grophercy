package grifts

import (
	"github.com/adilw3nomad/grophercy/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
