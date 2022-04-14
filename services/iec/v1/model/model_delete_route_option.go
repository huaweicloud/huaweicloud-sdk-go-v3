package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除路由对象
type DeleteRouteOption struct {
	// 功能说明：路由的目的网段  约束：合法的CIDR格式

	Destination string `json:"destination"`
}

func (o DeleteRouteOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRouteOption struct{}"
	}

	return strings.Join([]string{"DeleteRouteOption", string(data)}, " ")
}
