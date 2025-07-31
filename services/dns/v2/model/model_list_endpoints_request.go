package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointsRequest Request Object
type ListEndpointsRequest struct {

	// 终端节点方向。 取值： inbound，表示入站终端节点。 outbound，表示出站终端节点。
	Direction string `json:"direction"`

	// 待查询的终端节点所属VPC的ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 终端节点名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 分页查询时配置每页返回的资源个数。 **约束限制：** 不涉及。 **取值范围：** 0~500。 **默认取值：** 500
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。 **约束限制：** 当设置marker不为空时，以marker为分页起始标识，offset不生效。。 **取值范围：** 0~2147483647。 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointsRequest", string(data)}, " ")
}
