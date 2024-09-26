package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseRouterProjectId 企业路由器的项目ID。
type EnterpriseRouterProjectId struct {

	// 企业路由器的项目ID。
	EnterpriseRouterProjectId string `json:"enterprise_router_project_id"`
}

func (o EnterpriseRouterProjectId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseRouterProjectId struct{}"
	}

	return strings.Join([]string{"EnterpriseRouterProjectId", string(data)}, " ")
}
