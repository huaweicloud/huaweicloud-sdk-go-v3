package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCompliancePackagesResponse Response Object
type SearchCompliancePackagesResponse struct {

	// 遵从包总数
	Count *int32 `json:"count,omitempty"`

	// 内置遵从包数量
	BuiltinComplianceNum *int32 `json:"builtin_compliance_num,omitempty"`

	// 自定义遵从包数量
	CustomizedComplianceNum *int32 `json:"customized_compliance_num,omitempty"`

	// 停用遵从包数量
	DisabledComplianceNum *int32 `json:"disabled_compliance_num,omitempty"`

	// 启用遵从包数量
	EnabledComplianceNum *int32 `json:"enabled_compliance_num,omitempty"`

	// 遵从包列表
	CompliancePackages *[]CompliancePackageModel `json:"compliance_packages,omitempty"`
	HttpStatusCode     int                       `json:"-"`
}

func (o SearchCompliancePackagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCompliancePackagesResponse struct{}"
	}

	return strings.Join([]string{"SearchCompliancePackagesResponse", string(data)}, " ")
}
