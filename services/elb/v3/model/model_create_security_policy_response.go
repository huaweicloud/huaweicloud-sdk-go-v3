package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateSecurityPolicyResponse struct {
	SecurityPolicy *SecurityPolicy `json:"security_policy,omitempty"`
	// 请求ID。  注：自动生成 。

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSecurityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityPolicyResponse", string(data)}, " ")
}
