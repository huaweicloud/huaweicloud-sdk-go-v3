package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentConfigurationRequest Request Object
type CreateComponentConfigurationRequest struct {

	// 组件id。
	ComponentId string `json:"component_id"`

	// 环境id。
	XEnvironmentID string `json:"X-Environment-ID"`

	// 租户的企业项目id。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 应用id。
	ApplicationId string `json:"application_id"`

	Body *CreateComponentConfigurationRequestBody `json:"body,omitempty"`
}

func (o CreateComponentConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CreateComponentConfigurationRequest", string(data)}, " ")
}
