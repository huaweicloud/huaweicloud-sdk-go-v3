package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePolicyRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreatePolicyRequestBody `json:"body,omitempty"`
}

func (o CreatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyRequest struct{}"
	}

	return strings.Join([]string{"CreatePolicyRequest", string(data)}, " ")
}
