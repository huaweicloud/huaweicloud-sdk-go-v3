package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskStream 作业流参数
type TaskStream struct {

	// 作业参数配置
	Common *interface{} `json:"common,omitempty"`

	Input *TaskInput `json:"input"`

	// 输出详情
	Outputs []TaskOutputs `json:"outputs"`
}

func (o TaskStream) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskStream struct{}"
	}

	return strings.Join([]string{"TaskStream", string(data)}, " ")
}
