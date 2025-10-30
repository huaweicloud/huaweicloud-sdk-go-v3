package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HoneyForensicInfo 蜜罐相关取证信息
type HoneyForensicInfo struct {

	// **参数解释**： 攻击源 IP **取值范围**： 不涉及
	AttackIp *string `json:"attack_ip,omitempty"`

	// **参数解释**： 隔离沙箱名称 **取值范围**： 不涉及
	SandboxName *string `json:"sandbox_name,omitempty"`

	// **参数解释**： 欺骗服务 **取值范围**： 不涉及
	ServiceName *string `json:"service_name,omitempty"`

	// **参数解释**： 攻击类型 **取值范围**： - probe：探测 - invade：入侵
	AttackType *string `json:"attack_type,omitempty"`

	// **参数解释**： 攻击手法 **取值范围**： 不涉及
	AttackMethodDesc *string `json:"attack_method_desc,omitempty"`

	// **参数解释**： 攻击行为 **取值范围**： 不涉及
	AttackDesc *string `json:"attack_desc,omitempty"`
}

func (o HoneyForensicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HoneyForensicInfo struct{}"
	}

	return strings.Join([]string{"HoneyForensicInfo", string(data)}, " ")
}
