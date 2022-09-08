package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteSessionRequest struct {

	// 机器人标识符。
	QabotId string `json:"qabot_id"`

	// 会话标识符。
	SessionId string `json:"session_id"`

	Body *PostQaSessionReq `json:"body,omitempty"`
}

func (o ExecuteSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSessionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteSessionRequest", string(data)}, " ")
}
