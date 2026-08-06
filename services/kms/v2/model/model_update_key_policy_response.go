package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeyPolicyResponse Response Object
type UpdateKeyPolicyResponse struct {

	// **参数解释：** 密钥策略ID **取值范围：** 不涉及
	PolicyId *string `json:"policy_id,omitempty"`

	Policy *UpdateKeyPolicyResponseBodyPolicy `json:"policy,omitempty"`

	// **参数解释：** 密钥策略描述信息 **取值范围：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 密钥策略最近更新时间 **取值范围：** 不涉及
	LastModifyTime *string `json:"last_modify_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateKeyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateKeyPolicyResponse", string(data)}, " ")
}
