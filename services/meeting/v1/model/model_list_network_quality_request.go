package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworkQualityRequest Request Object
type ListNetworkQualityRequest struct {

	// 会议id
	Conferenceid string `json:"conferenceid"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
	ConfToken string `json:"confToken"`

	// 会议的appId
	Appid *string `json:"appid,omitempty"`

	// 会议UUID，MMR房间ID
	Confuuid string `json:"confuuid"`

	Body *RestQosRequestDto `json:"body,omitempty"`
}

func (o ListNetworkQualityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkQualityRequest struct{}"
	}

	return strings.Join([]string{"ListNetworkQualityRequest", string(data)}, " ")
}
