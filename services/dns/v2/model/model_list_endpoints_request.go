package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEndpointsRequest struct {

	// 待查询的endpoint的方向。 取值：inbound，表示入站规则。
	Direction string `json:"direction"`

	// 待查询的endpoint所属vpc的id。
	VpcId *string `json:"vpc_id,omitempty"`

	// 每页返回的资源个数。 取值范围：0~500，取值一般为10，20，50。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始页码，起始值为0。 当前设置marker不为空时，以marker为分页起始标识。取值范围：0~2147483647。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointsRequest", string(data)}, " ")
}
