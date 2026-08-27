package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantUpgradeStrategiesResponse Response Object
type ListTenantUpgradeStrategiesResponse struct {

	// 总数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 策略列表
	Strategies     *[]TenantUpgradeStrategyDetailRsp `json:"strategies,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListTenantUpgradeStrategiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantUpgradeStrategiesResponse struct{}"
	}

	return strings.Join([]string{"ListTenantUpgradeStrategiesResponse", string(data)}, " ")
}
