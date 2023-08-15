package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostRouteResponse Response Object
type ListHostRouteResponse struct {

	// 路由线路数量
	Total *int32 `json:"total,omitempty"`

	// 单个路由线路详细信息
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
