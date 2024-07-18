package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnAccessPolicyResponse Response Object
type ShowVpnAccessPolicyResponse struct {
	AccessPolicy *VpnAccessPolicy `json:"access_policy,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowVpnAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowVpnAccessPolicyResponse", string(data)}, " ")
}
