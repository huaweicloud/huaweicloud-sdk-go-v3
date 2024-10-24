package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlatformLiveInfo 平台直播信息
type PlatformLiveInfo struct {

	// 直播平台ID。
	PlatformId string `json:"platform_id"`

	// 直播ID。如果配置，则段落切换回调中会携带该信息。只能包含英文、数字、-、_。 美团对应liveId
	LiveId string `json:"live_id"`
}

func (o PlatformLiveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlatformLiveInfo struct{}"
	}

	return strings.Join([]string{"PlatformLiveInfo", string(data)}, " ")
}
