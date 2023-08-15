package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperateHistorySession struct {

	// 会话时长，格式：hh:ii:ss
	Duration *string `json:"duration,omitempty"`

	// 会话id
	SessionId *int64 `json:"session_id,omitempty"`

	// 会话开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 会话结束时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o OperateHistorySession) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateHistorySession struct{}"
	}

	return strings.Join([]string{"OperateHistorySession", string(data)}, " ")
}
