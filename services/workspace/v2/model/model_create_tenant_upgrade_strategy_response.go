package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTenantUpgradeStrategyResponse Response Object
type CreateTenantUpgradeStrategyResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateTenantUpgradeStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTenantUpgradeStrategyResponse struct{}"
	}

	return strings.Join([]string{"CreateTenantUpgradeStrategyResponse", string(data)}, " ")
}
