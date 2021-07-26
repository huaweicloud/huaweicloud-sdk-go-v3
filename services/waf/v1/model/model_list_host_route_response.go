package model

import (
	"encoding/json"

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
	data, err := json.Marshal(o)
	if err != nil {
		return "ListHostRouteResponse struct{}"
	}

	return strings.Join([]string{"ListHostRouteResponse", string(data)}, " ")
}
