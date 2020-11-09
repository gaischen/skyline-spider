package pili

import (
	"github.com/pili-engineering/pili-sdk-go.v2/pili"
	"github.com/vanga-top/skyline-spider/api_server/config"
)

func CreateRTMPPushURL(config config.Config) (string, error) {
	streamKeyPrefix := "sdkvanga_ls1604365190496857000"
	mac := &pili.MAC{AccessKey: config.AccessKey, SecretKey: []byte(config.SecretKey)}
	client := pili.New(mac, nil)
	hub := client.Hub(config.HubName)

	stream := hub.Stream(streamKeyPrefix)
	if stream != nil {
		url := genRTMPublishURL(config, streamKeyPrefix, mac)
		return url, nil
	}

	_, err := hub.Create(streamKeyPrefix)
	if err != nil {
		return "", err
	}
	url := genRTMPublishURL(config, streamKeyPrefix, mac)
	return url, nil
}

func genRTMPublishURL(config config.Config, streamKey string, mac *pili.MAC) string {
	return pili.RTMPPublishURL("pili-publish.vanga.top", config.HubName, streamKey, mac, 3600)
}

/**
-- get rtmp play url
*/
func GetRTMPPlayURL(config config.Config, streamKey string) string {
	return pili.RTMPPlayURL("pili-live-rtmp.vanga.top", config.HubName, streamKey)
}
