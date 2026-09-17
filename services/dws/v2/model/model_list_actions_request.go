package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListActionsRequest Request Object
type ListActionsRequest struct {

	// **参数解释**： 分页查询，每页大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActionsRequest struct{}"
	}

	return strings.Join([]string{"ListActionsRequest", string(data)}, " ")
}
