package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAndReviseRequest Request Object
type UpdateAndReviseRequest struct {

	// **参数解释：**  应用唯一标识。  **约束限制：**  不涉及。  **取值范围：**  由英文字母和数字组成，且长度为32个字符。  **默认取值：**  不涉及。
	Identifier string `json:"identifier"`

	// **参数解释：**  数据模型的英文名称。  **约束限制：**  不涉及。  **取值范围：**  大写字母开头，只能包含字母、数字、\"_\"，且长度为[1-60]个字符。  **默认取值：**  不涉及。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionReviseAndUpdateDtoVersionModel `json:"body,omitempty"`
}

func (o UpdateAndReviseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAndReviseRequest struct{}"
	}

	return strings.Join([]string{"UpdateAndReviseRequest", string(data)}, " ")
}
