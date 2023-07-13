package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRouteTableReq
type UpdateRouteTableReq struct {

	// 功能说明：路由表名称  取值范围：0-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 功能说明：路由表描述信息  取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`

	Routes *RouteTableRouteAction `json:"routes,omitempty"`
}

func (o UpdateRouteTableReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteTableReq struct{}"
	}

	return strings.Join([]string{"UpdateRouteTableReq", string(data)}, " ")
}
