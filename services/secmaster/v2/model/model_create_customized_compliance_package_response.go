package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomizedCompliancePackageResponse Response Object
type CreateCustomizedCompliancePackageResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误描述
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCustomizedCompliancePackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomizedCompliancePackageResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomizedCompliancePackageResponse", string(data)}, " ")
}
