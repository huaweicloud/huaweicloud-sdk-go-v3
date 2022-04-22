package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTbSessionRequest struct {

	// 话务机器人ID
	BotId string `json:"bot_id"`
}

func (o CreateTbSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTbSessionRequest struct{}"
	}

	return strings.Join([]string{"CreateTbSessionRequest", string(data)}, " ")
}
