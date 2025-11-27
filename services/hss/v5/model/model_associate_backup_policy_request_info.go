package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateBackupPolicyRequestInfo struct {

	// **参数解释**: 需要绑定的存储库id **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	VaultId *string `json:"vault_id,omitempty"`

	// **参数解释**: 需要绑定的策略id **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	PolicyId *string `json:"policy_id,omitempty"`
}

func (o AssociateBackupPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateBackupPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"AssociateBackupPolicyRequestInfo", string(data)}, " ")
}
