package pili

import (
	"fmt"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/rtc"
	"github.com/vanga-top/skyline-spider/api_server/config"
	"time"
)

func GetRoomToken(config config.Config) (string, error) {
	mac := auth.New(config.AccessKey, config.SecretKey)
	manager := rtc.NewManager(mac)

	app, err := manager.GetApp("f99wmp4jr")
	if err != nil {
		fmt.Println("get rtc app error", err)
		return "", err
	}

	token, err := manager.GetRoomToken(rtc.RoomAccess{AppID: app.AppID, RoomName: "testRtcRoom", UserID: "u123456789", ExpireAt: time.Now().Unix() + 3600})
	if err != nil {
		fmt.Println("get rtc room token error", err)
		return "", err
	}
	return token, nil
}
