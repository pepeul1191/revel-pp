package config

import (
	"github.com/revel/revel"
)

func BeforeAllFilter(c *revel.Controller, fc []revel.Filter) {
	c.Response.ContentType = "text/html; charset=UTF-8"
	c.Response.Out.Header().Add("Server", "Ubuntu")
	fc[0](c, fc[1:])
}
