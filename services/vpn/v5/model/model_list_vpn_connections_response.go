package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnConnectionsResponse Response Object
type ListVpnConnectionsResponse struct {
	VpnConnections *[]ResponseVpnConnection `json:"vpn_connections,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 租户下连接总数
	TotalCount *int64 `json:"total_count,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ListVpnConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListVpnConnectionsResponse", string(data)}, " ")
}
