package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HotQuestionInfo 热点问题信息。
type HotQuestionInfo struct {

	// 热点问题ID。
	HotQuestionId *string `json:"hot_question_id,omitempty"`

	// 热点问题。
	HotQuestion *string `json:"hot_question,omitempty"`

	// 机器人ID。
	RobotId *string `json:"robot_id,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o HotQuestionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotQuestionInfo struct{}"
	}

	return strings.Join([]string{"HotQuestionInfo", string(data)}, " ")
}
