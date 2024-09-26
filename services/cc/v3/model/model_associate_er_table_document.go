package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateErTableDocument 中心网络平面关联的ER路由表。
type AssociateErTableDocument struct {

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 企业路由器的ID。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// 企业路由器的路由表ID。
	EnterpriseRouterTableId string `json:"enterprise_router_table_id"`
}

func (o AssociateErTableDocument) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateErTableDocument struct{}"
	}

	return strings.Join([]string{"AssociateErTableDocument", string(data)}, " ")
}
