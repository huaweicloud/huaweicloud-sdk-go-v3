package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestMuteParticipantReqBody 静音与会者请求。
type RestMuteParticipantReqBody struct {

	// - 0: 取消静音 - 1: 静音
	IsMute int32 `json:"isMute"`
}

func (o RestMuteParticipantReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestMuteParticipantReqBody struct{}"
	}

	return strings.Join([]string{"RestMuteParticipantReqBody", string(data)}, " ")
}
