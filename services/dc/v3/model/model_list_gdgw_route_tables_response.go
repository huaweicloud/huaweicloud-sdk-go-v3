package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGdgwRouteTablesResponse Response Object
type ListGdgwRouteTablesResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 路由表详细对象
	GdgwRoutetable *[]ShowGdgwRoutetable `json:"gdgw_routetable,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGdgwRouteTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGdgwRouteTablesResponse struct{}"
	}

	return strings.Join([]string{"ListGdgwRouteTablesResponse", string(data)}, " ")
}
