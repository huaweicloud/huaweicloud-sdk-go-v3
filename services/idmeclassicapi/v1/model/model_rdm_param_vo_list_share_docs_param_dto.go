package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListShareDocsParamDto struct {

	// **参数解释**：  应用ID。  **约束限制**：  不涉及。  **取值范围**：  由英文字母和数字组成，且长度为32个字符。  **默认取值**：  不涉及。
	ApplicationId *string `json:"applicationId,omitempty"`

	// **参数解释：**  请求参数对象。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Params []ShareDocsParamDto `json:"params"`
}

func (o RdmParamVoListShareDocsParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListShareDocsParamDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListShareDocsParamDto", string(data)}, " ")
}
