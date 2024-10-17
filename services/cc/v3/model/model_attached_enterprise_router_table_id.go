package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachedEnterpriseRouterTableId 被挂载的企业路由器的路由表ID。
type AttachedEnterpriseRouterTableId struct {

	// 被挂载的企业路由器的路由表ID。
	AttachedErTableId string `json:"attached_er_table_id"`
}

func (o AttachedEnterpriseRouterTableId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachedEnterpriseRouterTableId struct{}"
	}

	return strings.Join([]string{"AttachedEnterpriseRouterTableId", string(data)}, " ")
}
