package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTenantUpgradeStrategyRequestBody 批量删除升级策略请求
type DeleteTenantUpgradeStrategyRequestBody struct {

	// 策略ID列表
	StrategyIds *[]string `json:"strategy_ids,omitempty"`
}

func (o DeleteTenantUpgradeStrategyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTenantUpgradeStrategyRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteTenantUpgradeStrategyRequestBody", string(data)}, " ")
}
