package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSessionRequest Request Object
type CreateSessionRequest struct {

	// 机器人标识符。
	QabotId string `json:"qabot_id"`
}

func (o CreateSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSessionRequest struct{}"
	}

	return strings.Join([]string{"CreateSessionRequest", string(data)}, " ")
}
