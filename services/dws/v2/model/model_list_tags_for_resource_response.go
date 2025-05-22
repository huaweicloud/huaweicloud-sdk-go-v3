package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsForResourceResponse Response Object
type ListTagsForResourceResponse struct {

	// **参数解释**： 标签列表。 **取值范围**： 不涉及。
	SysTags *[]Tag `json:"sys_tags,omitempty"`

	// **参数解释**： 标签数量。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTagsForResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsForResourceResponse struct{}"
	}

	return strings.Join([]string{"ListTagsForResourceResponse", string(data)}, " ")
}
