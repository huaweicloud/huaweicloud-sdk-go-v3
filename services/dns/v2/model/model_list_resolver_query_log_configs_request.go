package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResolverQueryLogConfigsRequest Request Object
type ListResolverQueryLogConfigsRequest struct {

	// **参数解释：** 分页查询时配置每页返回的资源个数。 **约束限制：** 不涉及。 **取值范围：** 0~500。 **默认取值：** 500
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页查询的起始资源ID。 - 查询第一页时，设置为空。 - 查询下一页时，设置为上一页最后一条资源的ID。  **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Marker *string `json:"marker,omitempty"`

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o ListResolverQueryLogConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResolverQueryLogConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListResolverQueryLogConfigsRequest", string(data)}, " ")
}
