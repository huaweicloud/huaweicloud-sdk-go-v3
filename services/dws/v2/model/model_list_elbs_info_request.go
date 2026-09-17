package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListElbsInfoRequest Request Object
type ListElbsInfoRequest struct {

	// **参数解释**： 虚拟私有云ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： 子网ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListElbsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListElbsInfoRequest struct{}"
	}

	return strings.Join([]string{"ListElbsInfoRequest", string(data)}, " ")
}
