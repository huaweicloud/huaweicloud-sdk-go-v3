package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartMeetingResponse Response Object
type StartMeetingResponse struct {

	// uuid。 > 废弃参数，请勿使用。
	Uuid *string `json:"uuid,omitempty"`

	// 会议所在区域的公网IP地址。
	RegionIP       *string `json:"regionIP,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartMeetingResponse struct{}"
	}

	return strings.Join([]string{"StartMeetingResponse", string(data)}, " ")
}
