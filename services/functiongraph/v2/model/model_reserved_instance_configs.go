package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReservedInstanceConfigs struct {

	// 函数URN
	FunctionUrn *string `json:"function_urn,omitempty"`

	// 限定类型, 支持version和alias
	QualifierType *string `json:"qualifier_type,omitempty"`

	// 限定类型对应的取值
	QualifierName *string `json:"qualifier_name,omitempty"`

	// 预留实例个数
	MinCount *int32 `json:"min_count,omitempty"`

	// 是否开启闲置模式配置
	IdleMode *bool `json:"idle_mode,omitempty"`

	TacticsConfig *TacticsConfig `json:"tactics_config,omitempty"`
}

func (o ReservedInstanceConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReservedInstanceConfigs struct{}"
	}

	return strings.Join([]string{"ReservedInstanceConfigs", string(data)}, " ")
}
