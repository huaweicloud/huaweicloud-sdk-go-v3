package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModelVersionCheckoutAndUpdateDtoVersionModel struct {

	// **参数解释**：  应用ID。  **约束限制**：  不涉及。  **取值范围**：  由英文字母和数字组成，且长度为32个字符。  **默认取值**：  不涉及。
	ApplicationId *string `json:"applicationId,omitempty"`

	// **参数解释：**  请求参数对象。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Params *[]VersionModelVersionCheckoutAndUpdateDtoVersionModel `json:"params,omitempty"`
}

func (o RdmParamVoListVersionModelVersionCheckoutAndUpdateDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModelVersionCheckoutAndUpdateDtoVersionModel struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModelVersionCheckoutAndUpdateDtoVersionModel", string(data)}, " ")
}
