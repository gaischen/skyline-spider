package main

import (
	"fmt"
	"github.com/pili-engineering/pili-sdk-go.v2/pili"
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

var streamKeyA = "sdkvanga_ls1604365190496857000"

//创建视频流
func TestCreateLiveStream(t *testing.T) {

	//streamKeyPrefix := "sdkvanga_ls" + strconv.FormatInt(time.Now().UnixNano(), 10)
	//streamKeyPrefix := "sdkvanga_ls1604365190496857000"
	// 初始化 client & hub.
	mac := &pili.MAC{AccessKey: AccessKey, SecretKey: []byte(SecretKey)}
	client := pili.New(mac, nil)
	hub := client.Hub(HubName)

	fmt.Println("创建流......")
	stream, err := hub.Create(streamKeyA)
	if err != nil {
		fmt.Println("创建流失败：", err)
		//return
	}
	info, err := stream.Info()
	if err != nil {
		fmt.Println("获取流信息失败：", err)
		return
	}
	fmt.Println(info)
}

func TestGetRTMPPushAddress(t *testing.T) {
	if AccessKey == "" {
		return
	}
	mac := &pili.MAC{AccessKey: AccessKey, SecretKey: []byte(SecretKey)}

	fmt.Println("RTMP 推流地址:")
	url := pili.RTMPPublishURL("pili-publish.vanga.top", HubName, streamKeyA, mac, 3600)
	fmt.Println(url)
}
