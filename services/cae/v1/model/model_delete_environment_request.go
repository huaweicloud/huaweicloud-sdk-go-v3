package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnvironmentRequest Request Object
type DeleteEnvironmentRequest struct {

	// 环境id。
	EnvironmentId string `json:"environment_id"`

	// 租户的企业项目id。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`
}

func (o DeleteEnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentRequest", string(data)}, " ")
}
