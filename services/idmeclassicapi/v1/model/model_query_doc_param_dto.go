package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryDocParamDto struct {

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  文档类型。  **约束限制**：  不涉及。  **取值范围**：  - directory: 目录 - pageDocument: Page文档 - boardDocument: Board文档 - mindDocument: Mind文档 - drawDocument: Draw文档  **默认取值**：  不涉及。
	Type *string `json:"type,omitempty"`
}

func (o QueryDocParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDocParamDto struct{}"
	}

	return strings.Join([]string{"QueryDocParamDto", string(data)}, " ")
}
