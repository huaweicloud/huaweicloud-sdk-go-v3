package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantUpgradeStrategyResponse Response Object
type UpdateTenantUpgradeStrategyResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateTenantUpgradeStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantUpgradeStrategyResponse struct{}"
	}

	return strings.Join([]string{"UpdateTenantUpgradeStrategyResponse", string(data)}, " ")
}
