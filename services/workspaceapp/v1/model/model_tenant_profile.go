package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantProfile 租户信息配置。
type TenantProfile struct {

	// 租户ID 同tenant_id。
	ProjectId *string `json:"project_id,omitempty"`

	// 租户名称。
	ProjectName *string `json:"project_name,omitempty"`

	// 租户的域ID。
	TenantDomainId *string `json:"tenant_domain_id,omitempty"`

	// 服务状态 * `active` - 激活 * `inactive` - 未激活
	ServiceStatus *string `json:"service_status,omitempty"`

	// 是否对接AD。 有AD的情况下，提示租户单会话模式和多会话模式都支持; 在没有AD的情况下，提示租户仅支持VDI单会话模式。
	OpenWithAd *bool `json:"open_with_ad,omitempty"`

	// 租户的域名称。
	TenantDomainName *string `json:"tenant_domain_name,omitempty"`

	// 租户信息创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o TenantProfile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantProfile struct{}"
	}

	return strings.Join([]string{"TenantProfile", string(data)}, " ")
}
