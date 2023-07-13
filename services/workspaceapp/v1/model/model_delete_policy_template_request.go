package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyTemplateRequest Request Object
type DeletePolicyTemplateRequest struct {

	// 策略模板id。
	PolicyTemplateId string `json:"policy_template_id"`
}

func (o DeletePolicyTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyTemplateRequest", string(data)}, " ")
}
