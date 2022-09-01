package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFunctionReservedInstancesCountRequestBody struct {

	// 预留实例个数
	Count int32 `json:"count" xml:"count"`

	// 是否开启闲置模式配置
	IdleMode *bool `json:"idle_mode,omitempty" xml:"idle_mode"`

	TacticsConfig *TacticsConfig `json:"tactics_config,omitempty" xml:"tactics_config"`
}

func (o UpdateFunctionReservedInstancesCountRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionReservedInstancesCountRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionReservedInstancesCountRequestBody", string(data)}, " ")
}
