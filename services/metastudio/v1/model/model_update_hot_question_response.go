package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHotQuestionResponse Response Object
type UpdateHotQuestionResponse struct {

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

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateHotQuestionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotQuestionResponse struct{}"
	}

	return strings.Join([]string{"UpdateHotQuestionResponse", string(data)}, " ")
}
