package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateL7policyRequestBody This is a auto create Body Object
type CreateL7policyRequestBody struct {
	L7policy *CreateL7policyReq `json:"l7policy"`
}

func (o CreateL7policyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7policyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateL7policyRequestBody", string(data)}, " ")
}
