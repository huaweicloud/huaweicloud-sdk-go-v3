package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskOutputs 作业具体输出数据
type TaskOutputs struct {

	// 输出数据
	Data *interface{} `json:"data"`

	// 输出类型
	Type string `json:"type"`
}

func (o TaskOutputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputs struct{}"
	}

	return strings.Join([]string{"TaskOutputs", string(data)}, " ")
}
