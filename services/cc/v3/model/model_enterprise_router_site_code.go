package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseRouterSiteCode 中心网络企业路由器的站点编码。
type EnterpriseRouterSiteCode struct {

	// 中心网络企业路由器的站点编码。
	EnterpriseRouterSiteCode string `json:"enterprise_router_site_code"`
}

func (o EnterpriseRouterSiteCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseRouterSiteCode struct{}"
	}

	return strings.Join([]string{"EnterpriseRouterSiteCode", string(data)}, " ")
}
