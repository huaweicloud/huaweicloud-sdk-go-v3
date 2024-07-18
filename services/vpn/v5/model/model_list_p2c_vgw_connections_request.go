package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListP2cVgwConnectionsRequest Request Object
type ListP2cVgwConnectionsRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`

	// 分页查询时每页返回的记录数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListP2cVgwConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListP2cVgwConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListP2cVgwConnectionsRequest", string(data)}, " ")
}
