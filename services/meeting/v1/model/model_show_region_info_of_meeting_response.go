package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRegionInfoOfMeetingResponse struct {
	// 会议所在区域的公网IP地址。

	RegionIP *string `json:"regionIP,omitempty"`
	// 会议所在区域的公网域名。

	RegionUrl      *string `json:"regionUrl,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRegionInfoOfMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRegionInfoOfMeetingResponse struct{}"
	}

	return strings.Join([]string{"ShowRegionInfoOfMeetingResponse", string(data)}, " ")
}
