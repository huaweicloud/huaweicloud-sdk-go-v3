package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDomainLoginPolicyResponse struct {
	LoginPolicy    *LoginPolicyResult `json:"login_policy,omitempty" xml:"login_policy"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateDomainLoginPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainLoginPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainLoginPolicyResponse", string(data)}, " ")
}
