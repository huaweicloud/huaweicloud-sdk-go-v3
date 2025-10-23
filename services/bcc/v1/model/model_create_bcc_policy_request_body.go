package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBccPolicyRequestBody CreateBccPolicyRequestBody
type CreateBccPolicyRequestBody struct {
	Policy *CreateBccPolicyRequestBodyPolicy `json:"policy"`
}

func (o CreateBccPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBccPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateBccPolicyRequestBody", string(data)}, " ")
}
