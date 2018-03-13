package controllers

import (
	"github.com/revel/revel"
	resty "gopkg.in/resty.v1"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	c.ViewArgs["CSSs"] = []string{
		"bower_components/bootstrap/dist/css/bootstrap.min",
		"bower_components/font-awesome/css/font-awesome.min",
	}
	c.ViewArgs["JSs"] = []string{
		"bower_components/jquery/dist/jquery.min",
		"bower_components/bootstrap/dist/js/bootstrap.min",
	}
	c.ViewArgs["Email"] = "astaxie@gmail.com"
	c.ViewArgs["HeadTitle"] = "Home"
	return c.RenderTemplate("home/index.html")
}

func (c App) Listar() revel.Result {
	resp, err := resty.R().Get("http://localhost:3000/departamento/listar")
	if err != nil {
		c.Response.Status = 500
		return c.RenderText("Error: No se puede conectar contra el servicio")
	} else {
		c.Response.ContentType = "text/html; charset=UTF-8"
		return c.RenderText(resp.String())
	}
}
