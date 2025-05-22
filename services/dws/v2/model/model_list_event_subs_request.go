package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventSubsRequest Request Object
type ListEventSubsRequest struct {

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 1000
	Limit *string `json:"limit,omitempty"`
}

func (o ListEventSubsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSubsRequest struct{}"
	}

	return strings.Join([]string{"ListEventSubsRequest", string(data)}, " ")
}
