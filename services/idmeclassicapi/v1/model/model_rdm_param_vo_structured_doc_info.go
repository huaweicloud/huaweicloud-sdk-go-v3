package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoStructuredDocInfo struct {

	// **参数解释**：  应用ID。  **约束限制**：  不涉及。  **取值范围**：  由英文字母和数字组成，且长度为32个字符。  **默认取值**：  不涉及。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *StructuredDocInfo `json:"params,omitempty"`
}

func (o RdmParamVoStructuredDocInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoStructuredDocInfo struct{}"
	}

	return strings.Join([]string{"RdmParamVoStructuredDocInfo", string(data)}, " ")
}
