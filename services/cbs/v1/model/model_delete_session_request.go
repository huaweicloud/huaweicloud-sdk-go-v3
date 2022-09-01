package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSessionRequest struct {

	// 机器人标识符。
	QabotId string `json:"qabot_id" xml:"qabot_id"`

	// 会话标识符。
	SessionId string `json:"session_id" xml:"session_id"`
}

func (o DeleteSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSessionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSessionRequest", string(data)}, " ")
}
