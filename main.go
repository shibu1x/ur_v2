package main

import "github.com/shibu1x/ur_v2/model"

func main() {
	model.ConnectDB()
	// model.GenerateModel()
	model.UpdateHousesAll()
}
