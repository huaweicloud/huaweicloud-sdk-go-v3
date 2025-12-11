package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRuleIdListRequestInfo 检查项ID列表
type CheckRuleIdListRequestInfo struct {

	// **参数解释**: 检查项ID列表 **约束限制**: 不涉及 **取值范围**: 元素个数范围0-2147483647 **默认取值**: 不涉及
	CheckRules *[]CheckRuleKeyInfoRequestInfo `json:"check_rules,omitempty"`
}

func (o CheckRuleIdListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleIdListRequestInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleIdListRequestInfo", string(data)}, " ")
}
