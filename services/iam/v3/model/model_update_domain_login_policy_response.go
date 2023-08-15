package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainLoginPolicyResponse Response Object
type UpdateDomainLoginPolicyResponse struct {
	LoginPolicy    *LoginPolicyResult `json:"login_policy,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateDomainLoginPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainLoginPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainLoginPolicyResponse", string(data)}, " ")
}
