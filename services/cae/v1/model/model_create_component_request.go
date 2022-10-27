package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateComponentRequest struct {

	// 环境id。
	XEnvironmentID string `json:"X-Environment-ID"`

	// 租户的企业项目id。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 应用id。
	ApplicationId string `json:"application_id"`

	Body *CreateComponentRequestBody `json:"body,omitempty"`
}

func (o CreateComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentRequest struct{}"
	}

	return strings.Join([]string{"CreateComponentRequest", string(data)}, " ")
}
