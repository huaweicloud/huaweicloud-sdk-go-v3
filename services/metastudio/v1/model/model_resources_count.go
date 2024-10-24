package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourcesCount 资源总量信息
type ResourcesCount struct {

	// 业务类型。
	BusinessType *string `json:"business_type,omitempty"`

	// 资源总数。
	Count float32 `json:"count,omitempty"`

	// 计费类型。 * PERIODIC: 包周期 * ONE_TIME：一次性 * ON_DEMAND：按需
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 资源来源。 * PURCHASED: 购买 * SP_ALLOCATED：SP分配 * ADMIN_ALLOCATED：系统管理员分配
	ResourceSource *string `json:"resource_source,omitempty"`
}

func (o ResourcesCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesCount struct{}"
	}

	return strings.Join([]string{"ResourcesCount", string(data)}, " ")
}
