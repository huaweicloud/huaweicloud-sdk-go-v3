package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelInfo 任务基本信息
type CancelInfo struct {

	// 执行记录id
	Id *string `json:"id,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`
}

func (o CancelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelInfo struct{}"
	}

	return strings.Join([]string{"CancelInfo", string(data)}, " ")
}
