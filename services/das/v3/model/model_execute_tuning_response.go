package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTuningResponse Response Object
type ExecuteTuningResponse struct {

	// SQL诊断消息ID列表
	MessageId *[]string `json:"message_id,omitempty"`

	// 诊断任务创建状态
	Status *bool `json:"status,omitempty"`

	// 是否超过诊断任务创建限额
	QuotaExceeded  *bool `json:"quota_exceeded,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ExecuteTuningResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTuningResponse struct{}"
	}

	return strings.Join([]string{"ExecuteTuningResponse", string(data)}, " ")
}
