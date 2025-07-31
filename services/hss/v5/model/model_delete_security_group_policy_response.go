package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityGroupPolicyResponse Response Object
type DeleteSecurityGroupPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityGroupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupPolicyResponse", string(data)}, " ")
}
