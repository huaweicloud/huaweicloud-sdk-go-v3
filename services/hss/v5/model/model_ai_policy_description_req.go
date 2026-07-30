package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyDescriptionReq **参数解释**： 策略描述 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
type AiPolicyDescriptionReq struct {
}

func (o AiPolicyDescriptionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyDescriptionReq struct{}"
	}

	return strings.Join([]string{"AiPolicyDescriptionReq", string(data)}, " ")
}
