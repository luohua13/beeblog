package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego"
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
	Category		string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Cteated         string `orm:"index"`
	Updated         string `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyCount      int64
	ReplyLastUserId int64
}

type Comment struct {
	Id 		int64
	Tid		int64
	Name	string
	Content	string     `orm:"size(1000)"`
	Created string
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic),new(Comment))
	orm.RegisterDriver(_SQLITE_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE_DRIVER, _DB_NAME, 10)
}

func IsDir(Dir string) bool {
	f, e := os.Stat(Dir)
	if e != nil {
		return false
	}

	return f.IsDir()
}

func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid,10,64)
	if err != nil {
		return  err
	}
	
	reply := &Comment{
		Tid:		tidNum,
		Name:		nickname,
		Content:	content,
		Created:	time.Now().Format("2006-01-02 15:04:05"),
	}
	
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		beego.Error(err)
	}
	return err
}

func DeleteReply(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	reply := &Comment{Id: cid}
	_, err = o.Delete(reply)
	return err
}

func GetAllReplies(tid string) (replies []*Comment,err error) {
	tidNum, err := strconv.ParseInt(tid,10,64)
	if err != nil {
		return nil, err
	}
	
	o := orm.NewOrm()

	replies = make([]*Comment, 0)

	qs := o.QueryTable("comment")

	_, err = qs.Filter("tid", tidNum).All(&replies)
	
	return replies, err
}

func ModifyTopic(tid, title, content, category string) error {
	o := orm.NewOrm()
	
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error(err)
		return err
	}

	topic := &Topic{Id: id}
	beego.Debug(topic)

	if o.Read(topic) == nil {
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now().Format("2006-01-02 15:04:05")
		o.Update(topic)
	}

	return err
}

func AddTopic(title, content, category string) error {
	o := orm.NewOrm()
	
	topic := &Topic {
		Title: title,
		Category: category,
		Content: content,
		Cteated: time.Now().Format("2006-01-02 15:04:05"),
		Updated: time.Now().Format("2006-01-02 15:04:05"),
	}
	
	_, err := o.Insert(topic)
	if err != nil {
		beego.Error(err)
	}
	return err
}

func GetTopicsByCategory(category string, IsDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error

	if IsDesc {
		_, err = qs.Filter("category", category).OrderBy("cteated").All(&topics)
	} else {
		_, err = qs.Filter("category", category).All(&topics)
	}
	return topics, err
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
	cate.TopicCount++ 
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

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	
	topics := make([]*Topic,0)
	qs := o.QueryTable("Topic")
	
	var err error
	if isDesc {
		_, err = qs.OrderBy("-cteated").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	
	return topics, err
}
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("Category")
	_, err := qs.All(&cates)

	return cates, err
}

func GetTopic(tid string) (*Topic,error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++

	_, err = o.Update(topic)
	return topic, err
}

func DeleteTopic(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: cid}
	_, err = o.Delete(topic)
	return err
}

func UpdateCategory(name string) error {
	o := orm.NewOrm()

	category := new(Category)

	qs := o.QueryTable(category)
	err := qs.Filter("title", name).One(category)
	if err != nil {
		return err
	}

	category.TopicCount++

	_, err = o.Update(category)
	return err

}

/* 检查是否分类存在 */

func CheckCategory(title string) bool {
	AllCategories, err := GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	for _, category := range AllCategories {
		if category.Title == title {
			return true
		}
	}
	return false
}
