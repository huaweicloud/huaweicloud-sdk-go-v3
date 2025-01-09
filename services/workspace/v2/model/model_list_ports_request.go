package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortsRequest Request Object
type ListPortsRequest struct {

	// 分页查询的起始资源ID。
	Marker *string `json:"marker,omitempty"`

	// 分页查询每页返回的记录个数。
	Limit *int32 `json:"limit,omitempty"`

	// ip地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 是否被使用。
	IsUsed *bool `json:"is_used,omitempty"`
}

func (o ListPortsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortsRequest struct{}"
	}

	return strings.Join([]string{"ListPortsRequest", string(data)}, " ")
}
