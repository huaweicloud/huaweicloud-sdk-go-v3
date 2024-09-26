package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateErInstanceDocument 中心网络实例详情。
type AssociateErInstanceDocument struct {

	// 企业路由器的ID。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// RegionID。
	RegionId string `json:"region_id"`
}

func (o AssociateErInstanceDocument) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateErInstanceDocument struct{}"
	}

	return strings.Join([]string{"AssociateErInstanceDocument", string(data)}, " ")
}
