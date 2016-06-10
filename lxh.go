package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

const (
	TYPE_TEXT = 1
	TYPE_IMG  = 2
)

type Joke struct {
	Id      int
	Content string
	Sort    float32
	Num     int
	Type    int
}

func (u *User) TableName() string {
	return "joke"
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8")
	orm.RegisterModel(new(Joke))
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	for i := 1; i <= 100; i++ {
		jokeID := fmt.Sprintf("%d", i)
		url := "http://lengxiaohua.com/joke/" + jokeID

		doc, err := goquery.NewDocument(url)
		if err != nil {
			fmt.Println(err)
			continue
		}

		body := doc.Find("div#joke_content_" + jokeID)
		content, _ := body.Html()
		types := TYPE_TEXT
		if strings.Contains(content, "<img") {
			types = TYPE_IMG
		} else {
			content = body.Text()
		}
		content = strings.TrimSpace(content)
		if len(content) > 0 {
			//fmt.Println(content)
			joke := new(Joke)
			joke.Content = content
			joke.Num = i
			joke.Type = types
			_, err = o.Insert(joke)
			if err != nil {
				fmt.Println(err)
			}
		}

		time.Sleep(10 * time.Millisecond)
	}
}
