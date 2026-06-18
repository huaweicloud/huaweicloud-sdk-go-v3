package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProviderDescription **参数解释**： 身份提供商描述。  **约束限制**： 长度范围为[0,255]。  **取值范围**： 不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\"、\"$\"、\"^\"和\"*\"的字符串。  **默认取值**： 不涉及。
type ProviderDescription struct {
}

func (o ProviderDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProviderDescription struct{}"
	}

	return strings.Join([]string{"ProviderDescription", string(data)}, " ")
}
