package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSecurityPolicyResponse struct {
	SecurityPolicy *SecurityPolicy `json:"security_policy,omitempty"`
	// 请求ID。  注：自动生成 。

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSecurityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityPolicyResponse", string(data)}, " ")
}
