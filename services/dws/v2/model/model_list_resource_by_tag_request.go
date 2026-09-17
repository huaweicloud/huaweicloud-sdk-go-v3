package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceByTagRequest Request Object
type ListResourceByTagRequest struct {

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 10
	Limit *string `json:"limit,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *string `json:"offset,omitempty"`

	Body *TagFilterRequestBody `json:"body,omitempty"`
}

func (o ListResourceByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceByTagRequest struct{}"
	}

	return strings.Join([]string{"ListResourceByTagRequest", string(data)}, " ")
}
