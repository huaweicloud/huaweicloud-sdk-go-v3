package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppWhitelistPolicyResponse Response Object
type CreateAppWhitelistPolicyResponse struct {

	// 策略ID
	PolicyId       *string `json:"policy_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppWhitelistPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppWhitelistPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateAppWhitelistPolicyResponse", string(data)}, " ")
}
