package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseRouterTableId 企业路由器的路由表ID。
type EnterpriseRouterTableId struct {

	// 企业路由器的路由表ID。
	EnterpriseRouterTableId string `json:"enterprise_router_table_id"`
}

func (o EnterpriseRouterTableId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseRouterTableId struct{}"
	}

	return strings.Join([]string{"EnterpriseRouterTableId", string(data)}, " ")
}
