package grifts

import (
  "github.com/gobuffalo/buffalo"
	"CookingBoyFrontPage/cooking_boy/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
