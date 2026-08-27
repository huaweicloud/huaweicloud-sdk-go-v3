package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantUpgradeStrategiesRequest Request Object
type ListTenantUpgradeStrategiesRequest struct {

	// 是否精确匹配名称
	IsAccurateName *bool `json:"is_accurate_name,omitempty"`

	// 策略名称（支持模糊查询）
	StrategyName *string `json:"strategy_name,omitempty"`

	// 策略类型：0-服务端 1-客户端
	StrategyType *int32 `json:"strategy_type,omitempty"`

	// 是否强制升级：0-否 1-是
	IsForceUpgrade *int32 `json:"is_force_upgrade,omitempty"`

	// 启用状态：0-禁用 1-启用
	Status *int32 `json:"status,omitempty"`

	// 协议策略优先级
	StrategyPriority *int32 `json:"strategy_priority,omitempty"`

	// 偏移量，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量，默认10，最大100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTenantUpgradeStrategiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantUpgradeStrategiesRequest struct{}"
	}

	return strings.Join([]string{"ListTenantUpgradeStrategiesRequest", string(data)}, " ")
}
