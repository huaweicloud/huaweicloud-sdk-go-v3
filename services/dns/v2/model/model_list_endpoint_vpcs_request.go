package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointVpcsRequest Request Object
type ListEndpointVpcsRequest struct {

	// 每页返回的资源个数。  取值范围：0~500  取值一般为10，20，50。默认值为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。  取值范围：0~2147483647  默认值为0。  当前设置marker不为空时，以marker为分页起始标识。
	Offset *int32 `json:"offset,omitempty"`

	// 待查询的VPC的ID。
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o ListEndpointVpcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointVpcsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointVpcsRequest", string(data)}, " ")
}
