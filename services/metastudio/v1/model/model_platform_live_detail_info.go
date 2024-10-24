package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlatformLiveDetailInfo 平台直播信息
type PlatformLiveDetailInfo struct {

	// 直播平台ID。
	PlatformId *string `json:"platform_id,omitempty"`

	// 直播平台。美团填写meituan
	Platform *string `json:"platform,omitempty"`

	// 授权账号信息。 美团平台对应：opBizCode
	Account *string `json:"account,omitempty"`

	// 直播ID。如果配置，则段落切换回调中会携带该信息。 美团对应liveId
	LiveId *string `json:"live_id,omitempty"`
}

func (o PlatformLiveDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlatformLiveDetailInfo struct{}"
	}

	return strings.Join([]string{"PlatformLiveDetailInfo", string(data)}, " ")
}
