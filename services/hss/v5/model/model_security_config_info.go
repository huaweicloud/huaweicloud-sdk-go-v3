package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityConfigInfo 配置检测信息
type SecurityConfigInfo struct {

	// **参数解释**： 风险级别 **取值范围**： - high：高危 - medium：中危 - low：低危 - safe：安全，无风险
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 基线名称 **取值范围**： 不涉及
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释**： 检查项数量 **取值范围**： 不涉及
	CheckRuleNum *int32 `json:"check_rule_num,omitempty"`

	// **参数解释**： 风险项数量 **取值范围**： 不涉及
	FailedRuleNum *int32 `json:"failed_rule_num,omitempty"`

	// **参数解释**： 最新检测时间 **取值范围**： 不涉及
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**： 基线描述信息 **取值范围**： 不涉及
	CheckTypeDesc *string `json:"check_type_desc,omitempty"`
}

func (o SecurityConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityConfigInfo struct{}"
	}

	return strings.Join([]string{"SecurityConfigInfo", string(data)}, " ")
}
