package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCceIntegrationProtectionRequest Request Object
type AddCceIntegrationProtectionRequest struct {

	// Region ID
	Region string `json:"region"`

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CceIntegrationProtectionRequestBody `json:"body,omitempty"`
}

func (o AddCceIntegrationProtectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCceIntegrationProtectionRequest struct{}"
	}

	return strings.Join([]string{"AddCceIntegrationProtectionRequest", string(data)}, " ")
}
