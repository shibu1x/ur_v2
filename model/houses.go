package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm/clause"
)

const TableNameHouse = "houses"

// House mapped from table <houses>
type House struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Code       string    `gorm:"column:code;not null" json:"code"`
	PrefCode   string    `gorm:"column:pref_code;not null" json:"pref_code"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	RoomsGotAt time.Time `gorm:"column:rooms_got_at;not null;default:1000-01-01 00:00:00" json:"rooms_got_at"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	Rooms      []Room    `gorm:"foreignKey:HouseCode;references:Code"`
}

// TableName House's table name
func (*House) TableName() string {
	return TableNameHouse
}

// UR APIが返す部屋のJSON構造
type JsonData []struct {
	PageIndex        string `json:"pageIndex"`
	RowMax           string `json:"rowMax"`
	RowMaxSp         string `json:"rowMaxSp"`
	RowMaxNext       string `json:"rowMaxNext"`
	PageMax          string `json:"pageMax"`
	AllCount         string `json:"allCount"`
	Block            string `json:"block"`
	Tdfk             string `json:"tdfk"`
	Shisya           string `json:"shisya"`
	Danchi           string `json:"danchi"`
	Shikibetu        string `json:"shikibetu"`
	FloorAll         string `json:"floorAll"`
	RoomDetailLink   string `json:"roomDetailLink"`
	RoomDetailLinkSp string `json:"roomDetailLinkSp"`
	System           []struct {
		IMG           string `json:"制度_IMG"`
		NAMING_FAILED string `json:"制度名"`
		HTML          string `json:"制度HTML"`
	} `json:"system"`
	Parking       any    `json:"parking"`
	Design        []any  `json:"design"`
	FeatureParam  []any  `json:"featureParam"`
	Traffic       any    `json:"traffic"`
	Place         any    `json:"place"`
	Kanris        any    `json:"kanris"`
	Kouzou        any    `json:"kouzou"`
	Soukosu       any    `json:"soukosu"`
	ID            string `json:"id"`
	Year          any    `json:"year"`
	Name          string `json:"name"`
	Shikikin      string `json:"shikikin"`
	Requirement   string `json:"requirement"`
	Madori        string `json:"madori"`
	Rent          string `json:"rent"`
	RentNormal    string `json:"rent_normal"`
	RentNormalCSS string `json:"rent_normal_css"`
	Commonfee     string `json:"commonfee"`
	CommonfeeSp   any    `json:"commonfee_sp"`
	Status        any    `json:"status"`
	Type          string `json:"type"`
	Floorspace    string `json:"floorspace"`
	Floor         string `json:"floor"`
	URLDetail     any    `json:"urlDetail"`
	URLDetailSp   any    `json:"urlDetail_sp"`
	Feature       any    `json:"feature"`
}

// 建物にある部屋の情報を取得する
func (h *House) UpdateRooms() {

	log.Println(h.Code)

	if time.Since(h.RoomsGotAt).Hours() < 12 {
		// 前回取得から一定時間経過していない
		return
	}

	now := time.Now()
	for i := 0; i < 100; i++ {
		data := h.getRooms(i)
		if len(data) == 0 {
			break
		}

		var rooms []Room
		for _, v := range data {
			data, _ := json.Marshal(v)
			rooms = append(rooms, Room{
				HouseCode: h.Code,
				RoomCode:  v.ID,
				Status:    "ready",
				Data:      string(data),
				GotAt:     now,
			})
		}

		// upsert
		DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "house_code"}, {Name: "room_code"}},     // key colume
			DoUpdates: clause.AssignmentColumns([]string{"status", "data", "got_at"}), // column needed to be updated
		}).Create(&rooms)

		pageMax, _ := strconv.Atoi(data[0].PageMax)
		if i >= pageMax-1 {
			break
		}
	}

	DB.Model(&h).Update("rooms_got_at", now)

	h.updateStatusClosed()
}

// 取得できなかったデータは受付終了とする
func (h *House) updateStatusClosed() {
	var rooms []Room
	DB.Model(&h).Where("status", "ready").Where("got_at < ?", time.Now().Format("2006-01-02")).Association("Rooms").Find(&rooms)
	for _, r := range rooms {
		DB.Model(&r).Update("status", "closed")
		fmt.Println("closed. " + r.RoomCode)
	}
}

// UR API から建物の部屋を取得する
func (h *House) getRooms(index int) JsonData {
	params := h.getSearchParams()
	for k, v := range map[string][]string{
		"orderByField": {"0"},
		"orderBySort":  {"0"},
		"pageIndex":    {strconv.Itoa(index)},
	} {
		params[k] = v
	}

	resp, err := http.Post(
		"https://chintai.r6.ur-net.go.jp/chintai/api/bukken/detail/detail_bukken_room/",
		"application/x-www-form-urlencoded",
		strings.NewReader(url.Values(params).Encode()),
	)

	if err != nil {
		log.Println(err)
		return nil
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data JsonData
	json.Unmarshal(body, &data)

	return data
}

// UR API に必要なパラメーターを取得する
func (h *House) getSearchParams() map[string][]string {
	return map[string][]string{
		"shisya":    {h.Code[:2]},
		"danchi":    {h.Code[3:6]},
		"shikibetu": {h.Code[len(h.Code)-1:]},
	}
}
