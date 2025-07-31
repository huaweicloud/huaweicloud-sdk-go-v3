package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityGroupPolicyResponse Response Object
type UpdateSecurityGroupPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSecurityGroupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupPolicyResponse", string(data)}, " ")
}
