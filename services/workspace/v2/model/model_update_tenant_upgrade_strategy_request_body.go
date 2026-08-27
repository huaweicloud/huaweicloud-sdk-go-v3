package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantUpgradeStrategyRequestBody 更新升级策略请求
type UpdateTenantUpgradeStrategyRequestBody struct {

	// 策略名称
	StrategyName *string `json:"strategy_name,omitempty"`

	// 是否强制升级：0-否 1-是
	IsForceUpgrade *int32 `json:"is_force_upgrade,omitempty"`

	// 低于此版本升级
	MinVersion *string `json:"min_version,omitempty"`

	// 升级目标版本
	TargetVersion *string `json:"target_version,omitempty"`

	// 策略描述
	StrategyDesc *string `json:"strategy_desc,omitempty"`

	// 优先级（数值越小优先级越高）
	StrategyPriority *int32 `json:"strategy_priority,omitempty"`

	// 状态：0-禁用 1-启用
	Status *int32 `json:"status,omitempty"`

	// 应用对象列表
	ApplyObjects *[]ApplyObjectInfo `json:"apply_objects,omitempty"`
}

func (o UpdateTenantUpgradeStrategyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantUpgradeStrategyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTenantUpgradeStrategyRequestBody", string(data)}, " ")
}
