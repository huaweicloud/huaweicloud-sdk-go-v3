package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowApplicationRequest struct {

	// 应用id。
	ApplicationId string `json:"application_id"`

	// 环境id。
	XEnvironmentID string `json:"X-Environment-ID"`

	// 租户的企业项目id。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`
}

func (o ShowApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationRequest struct{}"
	}

	return strings.Join([]string{"ShowApplicationRequest", string(data)}, " ")
}
