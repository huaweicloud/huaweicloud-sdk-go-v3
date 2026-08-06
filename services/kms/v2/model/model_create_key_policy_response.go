package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKeyPolicyResponse Response Object
type CreateKeyPolicyResponse struct {

	// **参数解释：** 密钥策略ID **取值范围：** 不涉及
	PolicyId       *string `json:"policy_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKeyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateKeyPolicyResponse", string(data)}, " ")
}
