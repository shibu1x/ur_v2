package main

import (
	"testing"
	"time"

	"github.com/shibu1x/ur_v2/model"
)

// 都道府県にある建物の情報を取得する
func TestUpdateHousesAll(t *testing.T) {
	model.ConnectDB()
	model.UpdateHousesAll()

	// 12 hours ago
	time := time.Now().Add(time.Duration(-12) * time.Hour)

	var houses []model.House
	model.DB.Where("rooms_got_at > ?", time.Format("2006-01-02 15:04:05")).Find(&houses)

	if len(houses) < 1 {
		t.Fatalf("Data not updated.")
	}
}
