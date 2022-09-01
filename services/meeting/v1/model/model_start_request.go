package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 开始会议请求参数
type StartRequest struct {

	// 会议ID
	ConferenceID string `json:"conferenceID" xml:"conferenceID"`

	// 主席密码
	Password string `json:"password" xml:"password"`
}

func (o StartRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartRequest struct{}"
	}

	return strings.Join([]string{"StartRequest", string(data)}, " ")
}
