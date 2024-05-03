package model

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/gocolly/colly"
	"gorm.io/gorm/clause"
)

const TableNamePref = "prefs"

// Pref mapped from table <prefs>
type Pref struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Code      string    `gorm:"column:code;not null" json:"code"`
	Region    string    `gorm:"column:region;not null" json:"region"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	IsCrawl   bool      `gorm:"column:is_crawl;not null" json:"is_crawl"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Houses    []House   `gorm:"foreignKey:PrefCode;references:Code"`
}

// TableName Pref's table name
func (*Pref) TableName() string {
	return TableNamePref
}

// 都道府県にある建物の情報を取得する
func UpdateHousesAll() {
	pattern = regexp.MustCompile(`\.\./(\d+_\d+)\.html`)

	var prefs []Pref
	DB.Where("is_crawl = ?", true).Find(&prefs)
	for _, v := range prefs {
		v.updateHouses()
		v.updateRooms()
	}
}

var pattern *regexp.Regexp

// 都道府県に紐ずく建物の情報を取得する
func (p *Pref) updateHouses() {

	c := colly.NewCollector(
		colly.AllowedDomains("www.ur-net.go.jp"),
	)

	var houses []House

	c.OnHTML("div.module_tables_apartment table tbody tr", func(e *colly.HTMLElement) {
		href := e.ChildAttr("a", "href")

		matches := pattern.FindStringSubmatch(href)
		if len(matches) == 0 {
			fmt.Println("Did not match. " + href)
			return
		}
		log.Println(matches[1])

		var house House
		err := DB.Where("code = ?", matches[1]).First(&house).Error
		if err == nil {
			return
		}
		houses = append(houses, House{Code: matches[1], PrefCode: p.Code, Name: e.ChildText("span.js-bukken-name")})
	})

	c.Visit("https://www.ur-net.go.jp/chintai/" + p.Region + "/" + p.Code + "/list/")

	// upsert
	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "code"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"pref_code": "name"}),
	}).Create(&houses)
}

// 都道府県に紐ずく建物の部屋情報を取得する
func (p *Pref) updateRooms() {
	var houses []House
	DB.Model(&p).Association("Houses").Find(&houses)

	for _, v := range houses {
		v.UpdateRooms()
	}
}
