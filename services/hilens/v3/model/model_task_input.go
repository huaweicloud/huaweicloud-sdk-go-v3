package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskInput 作业具体输入数据
type TaskInput struct {

	// 输入数据
	Data *interface{} `json:"data"`

	// 输入类型
	Type string `json:"type"`
}

func (o TaskInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInput struct{}"
	}

	return strings.Join([]string{"TaskInput", string(data)}, " ")
}
