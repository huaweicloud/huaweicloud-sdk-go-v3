package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FederationId **参数解释**：联邦ID。 **约束限制**：如果作业未提交到联邦端点，该字段为空字符串。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
type FederationId struct {
}

func (o FederationId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FederationId struct{}"
	}

	return strings.Join([]string{"FederationId", string(data)}, " ")
}
