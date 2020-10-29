package main

import (
	"fmt"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/rtc"
	"testing"
	"time"
)

var manager *rtc.Manager

func TestRtcRoomToken(t *testing.T) {
	mac := auth.New(AccessKey, SecretKey)
	manager = rtc.NewManager(mac)
	var (
		appID    string = "f99wmp4jr"
		roomName string = "test_room_1"
		userID   string = "vanga"
	)

	app, err := manager.CreateApp(rtc.AppInitConf{
		Title:          "vanga_app_info",
		MaxUsers:       1000,
		NoAutoKickUser: false,
	})

	if err != nil {
		fmt.Println(err)
		panic("error..create app")
	}

	appID = app.AppID

	token, err := manager.GetRoomToken(rtc.RoomAccess{AppID: appID, RoomName: roomName, UserID: userID, ExpireAt: time.Now().Unix() + 3600})
	if err != nil {
		fmt.Println(err)
		panic("error")
	}

	fmt.Println("appId:", app.AppID)
	fmt.Println("roomToken:", token)
	fmt.Println(manager.ListUser(appID, roomName))
}
