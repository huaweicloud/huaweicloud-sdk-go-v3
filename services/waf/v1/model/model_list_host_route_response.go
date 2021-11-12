package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHostRouteResponse struct {
	// 数量

	Total *int32 `json:"total,omitempty"`
	// 路由信息body

	Items          *[]RouteBody `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListHostRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostRouteResponse struct{}"
	}

	return strings.Join([]string{"ListHostRouteResponse", string(data)}, " ")
}
