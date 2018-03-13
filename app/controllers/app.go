package controllers

import (
	"github.com/revel/revel"
	resty "gopkg.in/resty.v1"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Listar() revel.Result {
	resp, err := resty.R().Get("http://localhost:3000/departamento/listar")
	if err != nil {
		c.Response.Status = 500
		return c.RenderText("Error: No se puede conectar contra el servicio")
	} else {
		return c.RenderText(resp.String())
	}
}
