package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnvironmentRequest Request Object
type CreateEnvironmentRequest struct {

	// 租户的企业项目id。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	Body *CreateEnvironmentRequestBody `json:"body,omitempty"`
}

func (o CreateEnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentRequest", string(data)}, " ")
}
