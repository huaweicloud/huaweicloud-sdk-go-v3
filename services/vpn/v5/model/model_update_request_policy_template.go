package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRequestPolicyTemplate struct {
	IkePolicy *UpdateVgwIkePolicy `json:"ike_policy,omitempty"`

	IpsecPolicy *UpdateVgwIpsecPolicy `json:"ipsec_policy,omitempty"`
}

func (o UpdateRequestPolicyTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRequestPolicyTemplate struct{}"
	}

	return strings.Join([]string{"UpdateRequestPolicyTemplate", string(data)}, " ")
}
