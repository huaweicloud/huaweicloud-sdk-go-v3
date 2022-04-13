package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteTbSessionRequest struct {
	// 话务机器人ID。

	BotId string `json:"bot_id"`
	// 会话ID。

	SessionId string `json:"session_id"`
}

func (o DeleteTbSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTbSessionRequest struct{}"
	}

	return strings.Join([]string{"DeleteTbSessionRequest", string(data)}, " ")
}
