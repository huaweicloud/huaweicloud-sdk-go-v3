package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePremiumHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreatePremiumHostRequestBody `json:"body,omitempty"`
}

func (o CreatePremiumHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePremiumHostRequest struct{}"
	}

	return strings.Join([]string{"CreatePremiumHostRequest", string(data)}, " ")
}
