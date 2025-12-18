package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParaGroupCopy struct {

	// **参数解释**：  新参数组的名称。  **参数范围**：  不涉及。
	NewName string `json:"new_name"`

	// **参数解释**：  描述。  **参数范围**：  不涉及。
	Description *string `json:"description,omitempty"`
}

func (o ParaGroupCopy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParaGroupCopy struct{}"
	}

	return strings.Join([]string{"ParaGroupCopy", string(data)}, " ")
}
