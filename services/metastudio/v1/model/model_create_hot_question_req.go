package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHotQuestionReq 创建热点问题请求。
type CreateHotQuestionReq struct {

	// 机器人ID。
	RobotId string `json:"robot_id"`

	// 热点问题。
	HotQuestion string `json:"hot_question"`
}

func (o CreateHotQuestionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHotQuestionReq struct{}"
	}

	return strings.Join([]string{"CreateHotQuestionReq", string(data)}, " ")
}
