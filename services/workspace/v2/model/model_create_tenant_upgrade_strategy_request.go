package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTenantUpgradeStrategyRequest Request Object
type CreateTenantUpgradeStrategyRequest struct {
	Body *CreateTenantUpgradeStrategyRequestBody `json:"body,omitempty"`
}

func (o CreateTenantUpgradeStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTenantUpgradeStrategyRequest struct{}"
	}

	return strings.Join([]string{"CreateTenantUpgradeStrategyRequest", string(data)}, " ")
}
