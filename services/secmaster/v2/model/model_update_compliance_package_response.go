package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCompliancePackageResponse Response Object
type UpdateCompliancePackageResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误描述
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCompliancePackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCompliancePackageResponse struct{}"
	}

	return strings.Join([]string{"UpdateCompliancePackageResponse", string(data)}, " ")
}
