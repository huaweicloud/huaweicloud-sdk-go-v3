package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLayoutRequest Request Object
type DeleteLayoutRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 布局UUID。
	UuID string `json:"uuID"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`
}

func (o DeleteLayoutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLayoutRequest struct{}"
	}

	return strings.Join([]string{"DeleteLayoutRequest", string(data)}, " ")
}
