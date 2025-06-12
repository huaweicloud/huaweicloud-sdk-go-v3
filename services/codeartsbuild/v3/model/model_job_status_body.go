package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobStatusBody 任务状态返回体
type JobStatusBody struct {

	// 任务状态
	Status *string `json:"status,omitempty"`

	// 任务终止状态
	AbortStatus *string `json:"abort_status,omitempty"`

	// 标签
	Label *string `json:"label,omitempty"`
}

func (o JobStatusBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobStatusBody struct{}"
	}

	return strings.Join([]string{"JobStatusBody", string(data)}, " ")
}
