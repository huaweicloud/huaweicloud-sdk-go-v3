package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListChatRequest Request Object
type ListChatRequest struct {

	// **参数解释**： 限制量，单次查询总量。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,1000]。 **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，查询起始偏移。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,100000000]。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 排序规则。 **约束限制**： 不涉及 **取值范围**： - DESC：降序。 - ASC：升序。 **默认取值**： DESC
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**： 对话名称，支持模糊搜索。 **约束限制**： 不涉及 **取值范围**： 取值范围为[1-500]个字符。 **默认取值**： 不涉及
	Title *string `json:"title,omitempty"`
}

func (o ListChatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChatRequest struct{}"
	}

	return strings.Join([]string{"ListChatRequest", string(data)}, " ")
}
