package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRouteResponse Response Object
type UpdateInstanceRouteResponse struct {

	// 单个路由线路详细信息
	Body           *[]RouteBody `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateInstanceRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRouteResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRouteResponse", string(data)}, " ")
}
