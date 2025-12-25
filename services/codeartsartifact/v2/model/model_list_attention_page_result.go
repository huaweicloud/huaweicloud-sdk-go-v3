package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAttentionPageResult struct {

	// **参数解释**： 总记录数。 **取值范围**： 不涉及。
	TotalRecords *int32 `json:"totalRecords,omitempty"`

	// **参数解释**： 总页数。 **取值范围**： 不涉及。
	TotalPages *int32 `json:"totalPages,omitempty"`

	// **参数解释**： 关注列表。 **取值范围**： 不涉及。
	Data *[]ListAttentionResult `json:"data,omitempty"`
}

func (o ListAttentionPageResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttentionPageResult struct{}"
	}

	return strings.Join([]string{"ListAttentionPageResult", string(data)}, " ")
}
