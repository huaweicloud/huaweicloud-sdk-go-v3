package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMultiViewRequest Request Object
type DeleteMultiViewRequest struct {

	// **参数解释：**  数据模型的英文名称。  **约束限制：**  不涉及。  **取值范围：**  大写字母开头，只能包含字母、数字、“_”，且长度为[1-60]个字符。  **默认取值：**  不涉及。
	MvModelName string `json:"mvModelName"`

	// 应用ID。
	Identifier string `json:"identifier"`

	Body *RdmParamVoMultiViewModelMasterIdModifierDto `json:"body,omitempty"`
}

func (o DeleteMultiViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMultiViewRequest struct{}"
	}

	return strings.Join([]string{"DeleteMultiViewRequest", string(data)}, " ")
}
