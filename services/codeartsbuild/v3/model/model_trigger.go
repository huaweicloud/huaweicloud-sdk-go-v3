package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Trigger 定时任务触发器
type Trigger struct {

	// 自定义参数
	Parameters []ParameterItem `json:"parameters"`

	// 触发器类型
	Name string `json:"name"`
}

func (o Trigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Trigger struct{}"
	}

	return strings.Join([]string{"Trigger", string(data)}, " ")
}
