package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2AgreementRule **参数解释**：协议规则信息。
type V2AgreementRule struct {

	// **参数解释**：协议名称。 **取值范围**：长度为1~128的英文字母、数字、下划线和中划线的组合。
	AgreementName *string `json:"agreement_name,omitempty"`

	// **参数解释**：协议展示名称。 **取值范围**：由中文字符、英文字母、数字及特殊字符组成，长度不超过512个字符。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**：协议版本，例如v1.0.0。 **取值范围**：长度为1~128的英文字母、数字、下划线、中划线和小数点的组合。
	Version *string `json:"version,omitempty"`

	// **参数解释**：协议链接。 **取值范围**：符合URL格式规范，由英文字母、数字及URL合法特殊字符组成，长度不超过2000个字符。
	Url *string `json:"url,omitempty"`
}

func (o V2AgreementRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2AgreementRule struct{}"
	}

	return strings.Join([]string{"V2AgreementRule", string(data)}, " ")
}
