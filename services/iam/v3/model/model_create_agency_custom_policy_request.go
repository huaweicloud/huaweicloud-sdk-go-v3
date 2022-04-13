package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAgencyCustomPolicyRequest struct {
	Body *CreateAgencyCustomPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateAgencyCustomPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyCustomPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateAgencyCustomPolicyRequest", string(data)}, " ")
}
