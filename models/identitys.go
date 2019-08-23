package models

import "github.com/astaxie/beego/orm"

type Identitys struct {
	Id int `orm:"column(id);size(255);null"`
	Number   string    `orm:"column(number);size(255);null"`
	//Employees []*Employees `orm:"reverse(many)" json:"Employees,omitempty"`
}
func init(){
	orm.RegisterModel(new(Identitys))
}
func (t *Identitys) TableName() string {
	return "identitys"
}
