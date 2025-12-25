package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmPolicyId **参数解释**： 告警策略ID。 **约束限制**： 不涉及。 **取值范围**： 长度为[1,64]个字符，只能包含字母、数字、- **默认取值**： 不涉及。
type AlarmPolicyId struct {
}

func (o AlarmPolicyId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmPolicyId struct{}"
	}

	return strings.Join([]string{"AlarmPolicyId", string(data)}, " ")
}
