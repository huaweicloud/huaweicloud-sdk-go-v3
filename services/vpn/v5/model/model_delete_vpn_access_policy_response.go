package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnAccessPolicyResponse Response Object
type DeleteVpnAccessPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpnAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpnAccessPolicyResponse", string(data)}, " ")
}
