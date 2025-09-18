package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleSetStatusReq struct {

	// **参数解释**： 规则模板实例状态。 **约束限制**： 不涉及。 **取值范围**： - true：开。 - false：关。 **默认取值**： 不涉及。
	IsValid bool `json:"is_valid"`
}

func (o UpdateRuleSetStatusReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleSetStatusReq struct{}"
	}

	return strings.Join([]string{"UpdateRuleSetStatusReq", string(data)}, " ")
}
