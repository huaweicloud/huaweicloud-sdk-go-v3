package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFunctionReservedInstancesCountResponse Response Object
type UpdateFunctionReservedInstancesCountResponse struct {

	// 预留实例个数
	Count *int32 `json:"count,omitempty"`

	// 是否开启闲置模式配置
	IdleMode *bool `json:"idle_mode,omitempty"`

	TacticsConfig  *TacticsConfig `json:"tactics_config,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateFunctionReservedInstancesCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionReservedInstancesCountResponse struct{}"
	}

	return strings.Join([]string{"UpdateFunctionReservedInstancesCountResponse", string(data)}, " ")
}
