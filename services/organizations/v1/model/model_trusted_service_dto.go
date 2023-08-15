package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrustedServiceDto 包含可信服务详细信息的结构，可信服务表示已启用与组织集成的服务。
type TrustedServiceDto struct {

	// 可信服务的名称。
	ServicePrincipal string `json:"service_principal"`

	// 可信服务与组织集成的日期。
	EnabledAt *sdktime.SdkTime `json:"enabled_at"`
}

func (o TrustedServiceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrustedServiceDto struct{}"
	}

	return strings.Join([]string{"TrustedServiceDto", string(data)}, " ")
}
