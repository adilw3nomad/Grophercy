package grifts

import (
	"github.com/adilw3nomad/grophercy/api/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
