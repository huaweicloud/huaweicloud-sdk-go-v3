package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityGroupPolicyResponse Response Object
type CreateSecurityGroupPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSecurityGroupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupPolicyResponse", string(data)}, " ")
}
