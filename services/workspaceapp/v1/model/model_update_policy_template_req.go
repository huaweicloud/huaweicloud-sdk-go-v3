package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePolicyTemplateReq struct {
	PolicyGroup *PolicyTemplate `json:"policy_group"`
}

func (o UpdatePolicyTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyTemplateReq struct{}"
	}

	return strings.Join([]string{"UpdatePolicyTemplateReq", string(data)}, " ")
}
