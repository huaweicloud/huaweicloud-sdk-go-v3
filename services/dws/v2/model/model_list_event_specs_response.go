package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventSpecsResponse Response Object
type ListEventSpecsResponse struct {

	// **参数解释**： 事件配置总数。 **取值范围**： 大于等于0。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 事件配置列表。 **取值范围**： 不涉及。
	EventSpecs     *[]EventSpecResponse `json:"event_specs,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListEventSpecsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSpecsResponse struct{}"
	}

	return strings.Join([]string{"ListEventSpecsResponse", string(data)}, " ")
}
