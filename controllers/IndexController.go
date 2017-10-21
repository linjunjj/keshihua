package controllers

import (
	"github.com/astaxie/beego"
	"keshihua/models"
)

type IndexController struct {
	beego.Controller
}

//首页
func (c *IndexController)index(){
	
}
//搜索请求
func (c *IndexController)search(){
	bookname:=c.Input().Get("bookname");
	if len(bookname)>0 {
		 book:=models.FindBookSeach(bookname)
		c.Data["data"]=book
	}else {
	}
}
//跳转请求
func (c *IndexController)requestTodetail