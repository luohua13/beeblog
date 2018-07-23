package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"time"
	"fmt"
	"strconv"
)

const (
	_DB_NAME       = "data/beeblog.db"
	_SQLITE_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	//Created         time.Time `orm:"index"`
	//Views           int64     `orm:"index"`
	//TopicTime       time.Time
	TopicCount      int64
	//TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Cteated         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE_DRIVER, _DB_NAME, 10)
}

func AddTopic(title, content string) error {
	o := orm.NewOrm()
	
	topic := &Topic {
		Title: title,
		Content: content,
		Cteated: time.Now(),
		Updated: time.Now(),
	}
	
	_, err := o.Insert(topic)
		
	return err
}
func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title: name}

	qs := o.QueryTable("Category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		fmt.Println("luohua0000000")
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		fmt.Println("luohua333333333")
		return err
	}

	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetAllTopics() ([]*Topic, error) {
	o := orm.NewOrm()
	
	topics := make([]*Topic,0)
	qs := o.QueryTable("Topic")
	_, err := qs.All(&topics)
	
	return topics, err
}
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("Category")
	_, err := qs.All(&cates)

	return cates, err
}

