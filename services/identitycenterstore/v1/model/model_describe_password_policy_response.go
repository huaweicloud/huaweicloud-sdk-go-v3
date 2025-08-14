package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribePasswordPolicyResponse Response Object
type DescribePasswordPolicyResponse struct {
	PasswordPolicy *PasswordPolicyDto `json:"password_policy,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o DescribePasswordPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribePasswordPolicyResponse struct{}"
	}

	return strings.Join([]string{"DescribePasswordPolicyResponse", string(data)}, " ")
}
