package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnAccessPolicyResponse Response Object
type CreateVpnAccessPolicyResponse struct {
	AccessPolicy *CreateVpnAccessPolicyResponseBodyAccessPolicy `json:"access_policy,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateVpnAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateVpnAccessPolicyResponse", string(data)}, " ")
}
