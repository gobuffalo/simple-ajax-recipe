package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/simple-ajax-recipe/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
