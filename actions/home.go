package actions

import (
	"time"

	"github.com/gobuffalo/buffalo"
)

func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

func TrafficCop(c buffalo.Context) error {
	time.Sleep(500 * time.Millisecond)
	p, _ := c.Value("badge").(string)
	switch p {
	case "success", "warning":
		c.Set("badge", p)
	default:
		c.Set("badge", "danger")
	}
	c.Set("now", time.Now())
	return c.Render(200, r.JavaScript("traffic.js"))
}
