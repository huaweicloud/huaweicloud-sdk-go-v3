package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRoutesResponse Response Object
type ListRoutesResponse struct {

	// 路由IP。
	RouteResps *[]RouteRespsResource `json:"routeResps,omitempty"`

	// 路由总数。
	TotalSize      *int32 `json:"totalSize,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoutesResponse struct{}"
	}

	return strings.Join([]string{"ListRoutesResponse", string(data)}, " ")
}
