package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachedEnterpriseRouterProjectId 被挂载的企业路由器的项目ID。
type AttachedEnterpriseRouterProjectId struct {

	// 被挂载的企业路由器的项目ID。
	AttachedErTableProjectId string `json:"attached_er_table_project_id"`
}

func (o AttachedEnterpriseRouterProjectId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachedEnterpriseRouterProjectId struct{}"
	}

	return strings.Join([]string{"AttachedEnterpriseRouterProjectId", string(data)}, " ")
}
