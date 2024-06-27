package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EchoTestPackageCheckResult struct {

	// 到期时间
	ExpirationTime *sdktime.SdkTime `json:"expiration_time,omitempty"`

	// 是否拥有license
	HasLicense *bool `json:"has_license,omitempty"`

	// 总量
	OriginalAmount *float64 `json:"original_amount,omitempty"`

	// 是否跨租户
	PackageUser *bool `json:"package_user,omitempty"`

	// 资源记录id
	ResourceId *string `json:"resource_id,omitempty"`

	// 套餐状态
	ResourceStatus *string `json:"resource_status,omitempty"`

	// 用量使用量
	ResourceUsed *float64 `json:"resource_used,omitempty"`

	// 套餐名称
	SpecCode *string `json:"spec_code,omitempty"`

	// 开始时间
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 租户id
	TenantId *string `json:"tenant_id,omitempty"`
}

func (o EchoTestPackageCheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EchoTestPackageCheckResult struct{}"
	}

	return strings.Join([]string{"EchoTestPackageCheckResult", string(data)}, " ")
}
