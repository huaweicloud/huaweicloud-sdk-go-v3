package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachedEnterpriseRouterRegionId ER路由器的regionID。
type AttachedEnterpriseRouterRegionId struct {

	// ER路由器的regionID。
	AttachedErTableRegionId string `json:"attached_er_table_region_id"`
}

func (o AttachedEnterpriseRouterRegionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachedEnterpriseRouterRegionId struct{}"
	}

	return strings.Join([]string{"AttachedEnterpriseRouterRegionId", string(data)}, " ")
}
