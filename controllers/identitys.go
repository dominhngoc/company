package controllers



import (

"github.com/astaxie/beego"
)
type IdentitysController struct {
	beego.Controller
}
func (c *IdentitysController) URLMapping() {
	c.Mapping("Get", c.Post)


}

