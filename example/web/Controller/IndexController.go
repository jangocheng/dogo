package Controller

import (
	"github.com/wuciyou/dogo"
	"github.com/wuciyou/dogo/config"
	"github.com/wuciyou/dogo/context"
)

type IndexController struct {
	Name string
	Age  int
}

func (index *IndexController) Index(c *context.Context) {

	name := make(map[string]string)
	name["name"] = "wuciyou"

	c.AddHeader("Server", "wuciyou")

	c.Assign(name)
	c.Display("View/index.html")

}

func (index *IndexController) ReturnJson(c *context.Context) {
	index.Name = "ReturnJson"
	index.Age = 25
	c.AjaxReturn(index)
}

func (index *IndexController) ReturnXml(c *context.Context) {
	index.Name = "ReturnXml"
	index.Age, _ = config.GetInt("LISTEN_PORT")
	c.AjaxReturn(index, "XML")
}
func (index *IndexController) ReturnAuto(c *context.Context) {
	index.Name = "Auto"
	index.Age = 25
	c.AjaxReturn(index)
}

func init() {
	index := &IndexController{}
	dogo.GetRouter("/", index.Index)
	dogo.GetRouter("/returnJson", index.ReturnJson)
	dogo.GetRouter("/returnXml", index.ReturnXml)
	dogo.GetRouter("/returnAuto", index.ReturnAuto)
}
