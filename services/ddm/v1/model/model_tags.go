package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tags struct {

	// **参数解释**：  标签名称。  **参数范围**：  不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**：  标签值。  **参数范围**：  不涉及。
	Value *string `json:"value,omitempty"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}
