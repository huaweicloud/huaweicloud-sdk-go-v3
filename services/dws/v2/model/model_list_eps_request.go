package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEpsRequest Request Object
type ListEpsRequest struct {

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页大小，默认10。 **约束限制**： 不涉及。 **取值范围**： 有效值大于等于1。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEpsRequest struct{}"
	}

	return strings.Join([]string{"ListEpsRequest", string(data)}, " ")
}
