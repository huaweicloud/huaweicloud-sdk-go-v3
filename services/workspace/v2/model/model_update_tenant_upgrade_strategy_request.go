package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantUpgradeStrategyRequest Request Object
type UpdateTenantUpgradeStrategyRequest struct {

	// 策略ID
	StrategyId string `json:"strategy_id"`

	Body *UpdateTenantUpgradeStrategyRequestBody `json:"body,omitempty"`
}

func (o UpdateTenantUpgradeStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantUpgradeStrategyRequest struct{}"
	}

	return strings.Join([]string{"UpdateTenantUpgradeStrategyRequest", string(data)}, " ")
}
