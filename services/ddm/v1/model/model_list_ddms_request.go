package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDdmsRequest Request Object
type ListDdmsRequest struct {

	// **参数解释**：  分页参数: 起始值。  **约束限制**：  不涉及。  **取值范围**：  大于等于0。  **默认取值**：  默认值是0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  分页参数: 每页记录数。  **约束限制**：  不涉及。  **取值范围**：  大于0且小于等于128。  **默认取值**：  默认值是10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDdmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDdmsRequest struct{}"
	}

	return strings.Join([]string{"ListDdmsRequest", string(data)}, " ")
}
