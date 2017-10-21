package models

import (
	"github.com/astaxie/beego/orm"
)
type Book struct {
	Id        int `orm:"pk;auto"`
	bookname  string `orm:"unique"`
	bookCode    string
	bookContent     string `orm:"unique"`
}
func FindBookByName(bookname string)(bool, Book){
   o :=orm.NewOrm()
	var book Book
	err := o.QueryTable(book).Filter("bookName", bookname).One(&book)
	return err !=orm.ErrNoRows,book
}
func FindBookSeach(bookname string) []*Book {
	var book  Book
	var books []*Book
	o:=orm.NewOrm()
	o.QueryTable(book).Filter("bookName",bookname).All(books)
	return books
}