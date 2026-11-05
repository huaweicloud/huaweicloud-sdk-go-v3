package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperinstanceGroupRequest Request Object
type ListHyperinstanceGroupRequest struct {

	// **参数解释**：每一页显示实例的数量。 **约束限制**：不涉及。 **取值范围**：[10,20,50]。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页记录的起始位置偏移量。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：排序的字段，多个字段使用(“,”)逗号分隔。 **约束限制**：不涉及。 **取值范围**：长度限制为128个字符，支持大小写字母、数字、中划线、下划线和逗号。 **默认取值**：不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：实例排序方式。 **约束限制**：不涉及。 **取值范围**： - ASC：升序 - DESC：降序  **默认取值**：DESC。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListHyperinstanceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperinstanceGroupRequest struct{}"
	}

	return strings.Join([]string{"ListHyperinstanceGroupRequest", string(data)}, " ")
}
