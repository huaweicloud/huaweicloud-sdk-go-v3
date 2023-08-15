package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityPolicyRequestBody This is a auto create Body Object
type CreateSecurityPolicyRequestBody struct {
	SecurityPolicy *CreateSecurityPolicyOption `json:"security_policy"`
}

func (o CreateSecurityPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecurityPolicyRequestBody", string(data)}, " ")
}
