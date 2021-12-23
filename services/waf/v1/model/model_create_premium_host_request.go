package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePremiumHostRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

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
