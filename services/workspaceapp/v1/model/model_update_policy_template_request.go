package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyTemplateRequest Request Object
type UpdatePolicyTemplateRequest struct {

	// 策略模板id。
	PolicyTemplateId string `json:"policy_template_id"`

	Body *UpdatePolicyTemplateReq `json:"body,omitempty"`
}

func (o UpdatePolicyTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyTemplateRequest", string(data)}, " ")
}
