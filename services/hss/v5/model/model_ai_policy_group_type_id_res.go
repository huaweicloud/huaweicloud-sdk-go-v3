package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyGroupTypeIdRes **参数解释**： 策略组ID **取值范围**： 最小值0，最大值2147483647
type AiPolicyGroupTypeIdRes struct {
}

func (o AiPolicyGroupTypeIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyGroupTypeIdRes struct{}"
	}

	return strings.Join([]string{"AiPolicyGroupTypeIdRes", string(data)}, " ")
}
