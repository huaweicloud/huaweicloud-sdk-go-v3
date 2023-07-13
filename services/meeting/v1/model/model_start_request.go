package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartRequest 激活会议请求。
type StartRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 会议密码。
	Password string `json:"password"`
}

func (o StartRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartRequest struct{}"
	}

	return strings.Join([]string{"StartRequest", string(data)}, " ")
}
