package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartMeetingResponse struct {
	// 会议主席鉴权uuid

	Uuid *string `json:"uuid,omitempty"`
	// 区域公网IP地址

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
