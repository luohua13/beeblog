package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	_DB_NAME       = "data/beeblog.db"
	_SQLITE_DRIVER = "sqlite3"
)

type Category struct {
	Id    int64
	Title string
	//Created         time.Time `orm:"index"`
	//Views           int64     `orm:"index"`
	//TopicTime       time.Time
	TopicCount int64
	//TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Labels          string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Cteated         string `orm:"index"`
	Updated         string `orm:"index"`
	Views           int64  `orm:"index"`
	Author          string `orm:"index"`
	ReplyCount      int64
	ReplyTime       string
	ReplyLastUserId int64
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string `orm:"size(1000)"`
	Created string `orm:"index"`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(Comment))
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
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now().Format("2006-01-02 15:04:05"),
	}

	o := orm.NewOrm()
	_, err = o.Insert(reply)

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

func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	replies = make([]*Comment, 0)

	qs := o.QueryTable("comment")

	_, err = qs.Filter("tid", tidNum).All(&replies)

	return replies, err
}

func ModifyTopic(tid, title, content, category, label, attachment string) error {
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()

	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error(err)
		return err
	}

	var oldCate, oldAttach string
	topic := &Topic{Id: id}

	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttach = topic.Attachment
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Attachment = attachment
		topic.Labels = label
		topic.Updated = time.Now().Format("2006-01-02 15:04:05")
		o.Update(topic)
	}

	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(cate)
		if err == nil {
			if cate.Title != category {
				cate.TopicCount--
				_, err = o.Update(cate)
				if err != nil {
					return err
				}
			}

		}
	}

	if len(oldAttach) > 0 {
		os.Remove(path.Join("attachment", oldAttach))
	}

	return err
}

func DeleteTopic(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: cid}
	var oldCate string
	if o.Read(topic) == nil {
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}

	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func AddTopic(title, content, category, label, attachment string) error {
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"
	o := orm.NewOrm()

	topic := &Topic{
		Title:      title,
		Category:   category,
		Content:    content,
		Attachment: attachment,
		Labels:     label,
		Cteated:    time.Now().Format("2006-01-02 15:04:05"),
		Updated:    time.Now().Format("2006-01-02 15:04:05"),
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

func AddCategory(name string, InitFlag bool) error {
	o := orm.NewOrm()

	cate := &Category{Title: name}

	qs := o.QueryTable("Category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	if !InitFlag {
		cate.TopicCount++
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func DelCategory(id string) error {
	category, err := GetCategory(id)
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")

	_, err = qs.Filter("category", category.Title).All(&topics)

	if err == nil {
		for _, topic := range topics {
			_, err = o.Delete(topic)
		}
	}

	cid, Err := strconv.ParseInt(id, 10, 64)
	if Err != nil {
		return Err
	}

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)

	return err

}

func GetAllTopics(cate string, label string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)
	qs := o.QueryTable("Topic")

	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("Category", cate)
		}
		if len(label) > 0 {
			qs = qs.Filter("labels__contains", "$"+label+"#")
		}
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

func GetCategory(tid string) (*Category, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	category := new(Category)

	qs := o.QueryTable("category")
	err = qs.Filter("id", tidNum).One(category)
	if err != nil {
		return nil, err
	}

	return category, err
}

func GetTopic(tid string) (*Topic, error) {
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

	topic.Labels = strings.Replace(strings.Replace(topic.Labels, "#", " ", -1), "$", "", -1)

	return topic, err
}

func UpdateCategory(name string, add bool) error {
	o := orm.NewOrm()

	category := new(Category)

	qs := o.QueryTable(category)
	err := qs.Filter("title", name).One(category)
	if err != nil {
		return err
	}

	if add {
		category.TopicCount++
	}
	_, err = o.Update(category)
	return err

}

func UpdateTopic(tid string, add bool) error {
	cid, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := new(Topic)
	qs := o.QueryTable(topic)
	err = qs.Filter("Id", cid).One(topic)

	if err != nil {
		return err
	}
	if add {
		topic.ReplyCount++
		topic.ReplyTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		replies := make([]*Comment, 0)
		qs = o.QueryTable("comment")
		_, err = qs.Filter("tid", cid).OrderBy("-created").All(&replies)
		if err != nil {
			return err
		}
		topic.ReplyCount--
		topic.ReplyTime = replies[0].Created
	}

	_, err = o.Update(topic)

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
