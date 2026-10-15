package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProvidedContextDto 可信上下文信息。
type ProvidedContextDto struct {

	// **参数解释**： 生成可信上下文断言的提供者的URN。  **约束限制**： 长度最多为1500个字符。格式为5段，用\":\"分隔，各段允许大小写字母、数字、\"+\"、\"/\"、\"=\"、\"-\"、\"_\"（第2段还额外允许 \"*\"），长度限制为：第1段至少1位，第2段0~255位，第3段1~64位，第4段1~64位，第5段至少1位（任意字符）。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	ContextProviderUrn string `json:"context_provider_urn"`

	// **参数解释**： 经过签名和加密的可信上下文断言。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	ContextAssertion string `json:"context_assertion"`
}

func (o ProvidedContextDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProvidedContextDto struct{}"
	}

	return strings.Join([]string{"ProvidedContextDto", string(data)}, " ")
}
