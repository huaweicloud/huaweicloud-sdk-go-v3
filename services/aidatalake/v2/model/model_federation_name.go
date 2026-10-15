package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FederationName **参数解释**：联邦名称。 **约束限制**：不涉及。 **取值范围**：长度为1~63个字符。 **默认取值**：不涉及。
type FederationName struct {
}

func (o FederationName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FederationName struct{}"
	}

	return strings.Join([]string{"FederationName", string(data)}, " ")
}
