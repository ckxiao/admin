package rbac

import (
	c "github.com/osgochina/admin/controllers"
	m "github.com/osgochina/admin/models/rbacmodels"
)

type GroupController struct {
	c.CommonController
}

func (this *GroupController) Index() {
	if this.IsAjax() {
		page, _ := this.GetInt("page")
		page_size, _ := this.GetInt("rows")
		sort := this.GetString("sort")
		order := this.GetString("order")
		if len(order) > 0 {
			if order == "desc" {
				sort = "-" + sort
			}
		} else {
			sort = "Id"
		}
		nodes, count := m.GetGrouplist(page, page_size, sort)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &nodes}
		this.ServeJson()
		return
	} else {
		this.TplNames = "easyui/rbac/group.tpl"
	}

}
func (this *GroupController) AddGroup() {
	g := m.Group{}
	if err := this.ParseForm(&g); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.AddGroup(&g)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *GroupController) UpdateGroup() {
	g := m.Group{}
	if err := this.ParseForm(&g); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.UpdateGroup(&g)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *GroupController) DelGroup() {
	Id, _ := this.GetInt("Id")
	status, err := m.DelGroupById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
