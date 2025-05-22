package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSecurityGroupRequestBody **参数解释**： 安全组请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ChangeSecurityGroupRequestBody struct {

	// **参数解释**： 安全组ID数组。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityGroups []string `json:"security_groups"`
}

func (o ChangeSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupRequestBody", string(data)}, " ")
}
