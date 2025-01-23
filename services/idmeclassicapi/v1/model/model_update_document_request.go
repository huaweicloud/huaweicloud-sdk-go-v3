package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDocumentRequest Request Object
type UpdateDocumentRequest struct {
	ModelName string `json:"modelName"`

	// **参数解释：**  应用唯一标识。  **约束限制：**  不涉及。  **取值范围：**  由英文字母和数字组成，且长度为32个字符。  **默认取值：**  不涉及。
	Identifier string `json:"identifier"`

	Body *RdmParamVoUpdateDocRequestDto `json:"body,omitempty"`
}

func (o UpdateDocumentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDocumentRequest struct{}"
	}

	return strings.Join([]string{"UpdateDocumentRequest", string(data)}, " ")
}
