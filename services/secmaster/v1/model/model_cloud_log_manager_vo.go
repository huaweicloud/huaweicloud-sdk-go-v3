package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudLogManagerVo struct {

	// 租户行管id
	TenantManagedDomainId *string `json:"tenant_managed_domain_id,omitempty"`

	// 平台行管id
	PlatformManagedDomainId *string `json:"platform_managed_domain_id,omitempty"`

	// regionId
	DwRegion *string `json:"dw_region,omitempty"`
}

func (o CloudLogManagerVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudLogManagerVo struct{}"
	}

	return strings.Join([]string{"CloudLogManagerVo", string(data)}, " ")
}
