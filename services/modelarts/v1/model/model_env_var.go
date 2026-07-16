package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvVar 精调训练环境变量信息
type EnvVar struct {

	// 标签
	Label *string `json:"label,omitempty"`

	// 描述信息
	Des *string `json:"des,omitempty"`

	// 环境变量名称
	EnvName *string `json:"env_name,omitempty"`

	// 环境变量类型
	EnvType *string `json:"env_type,omitempty"`

	// 环境变量值
	Value *string `json:"value,omitempty"`

	// 环境变量是否可修改
	Modifiable *bool `json:"modifiable,omitempty"`

	// 环境变量是否展示
	Displayable *bool `json:"displayable,omitempty"`

	// 环境变量使用阶段
	UsedSteps *[]string `json:"used_steps,omitempty"`
}

func (o EnvVar) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvVar struct{}"
	}

	return strings.Join([]string{"EnvVar", string(data)}, " ")
}
