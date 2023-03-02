package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务参数对象。
type TaskParam struct {

	// 参数名称。
	ParamName string `json:"param_name"`

	// 参数类型。
	ParamType *string `json:"param_type,omitempty"`

	// 参数分组。
	ParamGroup *string `json:"param_group,omitempty"`

	// 参数初始值。
	DefaultValue *string `json:"default_value,omitempty"`
}

func (o TaskParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskParam struct{}"
	}

	return strings.Join([]string{"TaskParam", string(data)}, " ")
}
