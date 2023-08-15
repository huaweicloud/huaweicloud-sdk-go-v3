package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainProtectPolicyResponse Response Object
type ShowDomainProtectPolicyResponse struct {
	ProtectPolicy  *ShowDomainProtectPolicyResponseBodyProtectPolicy `json:"protect_policy,omitempty"`
	HttpStatusCode int                                               `json:"-"`
}

func (o ShowDomainProtectPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainProtectPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainProtectPolicyResponse", string(data)}, " ")
}
