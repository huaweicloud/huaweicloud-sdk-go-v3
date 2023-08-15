package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainLoginPolicyRequestBody
type UpdateDomainLoginPolicyRequestBody struct {
	LoginPolicy *LoginPolicyOption `json:"login_policy"`
}

func (o UpdateDomainLoginPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainLoginPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainLoginPolicyRequestBody", string(data)}, " ")
}
