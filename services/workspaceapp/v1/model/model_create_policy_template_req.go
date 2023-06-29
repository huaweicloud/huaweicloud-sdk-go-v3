package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyTemplateReq 创建策略模板的请求
type CreatePolicyTemplateReq struct {
	PolicyGroup *PolicyTemplate `json:"policy_group"`
}

func (o CreatePolicyTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyTemplateReq struct{}"
	}

	return strings.Join([]string{"CreatePolicyTemplateReq", string(data)}, " ")
}
