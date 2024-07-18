package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnAccessPolicyResponse Response Object
type UpdateVpnAccessPolicyResponse struct {
	AccessPolicy *VpnAccessPolicy `json:"access_policy,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateVpnAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpnAccessPolicyResponse", string(data)}, " ")
}
