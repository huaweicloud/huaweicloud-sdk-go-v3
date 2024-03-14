package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DelegatedServiceDto 包含有关账号是可信服务委托管理员的信息。
type DelegatedServiceDto struct {

	// 服务主体的名称。
	ServicePrincipal string `json:"service_principal"`

	// 账号成为此服务的委托管理员的日期。
	DelegationEnabledAt *sdktime.SdkTime `json:"delegation_enabled_at"`
}

func (o DelegatedServiceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelegatedServiceDto struct{}"
	}

	return strings.Join([]string{"DelegatedServiceDto", string(data)}, " ")
}
