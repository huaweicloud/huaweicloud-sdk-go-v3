package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSecurityPolicyResponse struct {
	SecurityPolicy *SecurityPolicy `json:"security_policy,omitempty"`
	// 请求ID。  注：自动生成 。

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSecurityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolicyResponse", string(data)}, " ")
}
