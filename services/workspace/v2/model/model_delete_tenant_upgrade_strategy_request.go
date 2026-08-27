package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTenantUpgradeStrategyRequest Request Object
type DeleteTenantUpgradeStrategyRequest struct {
	Body *DeleteTenantUpgradeStrategyRequestBody `json:"body,omitempty"`
}

func (o DeleteTenantUpgradeStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTenantUpgradeStrategyRequest struct{}"
	}

	return strings.Join([]string{"DeleteTenantUpgradeStrategyRequest", string(data)}, " ")
}
