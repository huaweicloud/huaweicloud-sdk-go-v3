package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTenantUpgradeStrategyResponse Response Object
type DeleteTenantUpgradeStrategyResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteTenantUpgradeStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTenantUpgradeStrategyResponse struct{}"
	}

	return strings.Join([]string{"DeleteTenantUpgradeStrategyResponse", string(data)}, " ")
}
