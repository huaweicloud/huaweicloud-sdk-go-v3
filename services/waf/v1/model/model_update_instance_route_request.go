package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRouteRequest Request Object
type UpdateInstanceRouteRequest struct {

	// **参数解释：** 域名Id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	InstanceId string `json:"instance_id"`

	// 路由线路列表
	Body *[]RouteBody `json:"body,omitempty"`
}

func (o UpdateInstanceRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRouteRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRouteRequest", string(data)}, " ")
}
