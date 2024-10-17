package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachedEnterpriseRouterSiteCode 被挂载的企业路由器的站点编码。
type AttachedEnterpriseRouterSiteCode struct {

	// 被挂载的企业路由器的站点编码。
	AttachedErTableSiteCode string `json:"attached_er_table_site_code"`
}

func (o AttachedEnterpriseRouterSiteCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachedEnterpriseRouterSiteCode struct{}"
	}

	return strings.Join([]string{"AttachedEnterpriseRouterSiteCode", string(data)}, " ")
}
