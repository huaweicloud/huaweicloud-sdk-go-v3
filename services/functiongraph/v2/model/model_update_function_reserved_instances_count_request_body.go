package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFunctionReservedInstancesCountRequestBody struct {

	// 预留实例个数
	Count int32 `json:"count"`

	// 是否开启闲置模式配置
	IdleMode *bool `json:"idle_mode,omitempty"`

	TacticsConfig *TacticsConfig `json:"tactics_config,omitempty"`
}

func (o UpdateFunctionReservedInstancesCountRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionReservedInstancesCountRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionReservedInstancesCountRequestBody", string(data)}, " ")
}
