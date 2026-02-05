package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationCopyReqV3 请求体
type ConfigurationCopyReqV3 struct {

	// **参数解释**：  新参数组的名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	NewName string `json:"new_name"`

	// **参数解释**：  描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Description *string `json:"description,omitempty"`
}

func (o ConfigurationCopyReqV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationCopyReqV3 struct{}"
	}

	return strings.Join([]string{"ConfigurationCopyReqV3", string(data)}, " ")
}
