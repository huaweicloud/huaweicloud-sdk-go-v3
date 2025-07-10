package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortsRequest Request Object
type ListPortsRequest struct {

	// 分页查询的起始资源ID。
	Marker *string `json:"marker,omitempty"`

	// 功能说明：每页返回的个数。取值范围：1~2000。默认值：2000。
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
