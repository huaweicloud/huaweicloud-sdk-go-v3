package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCloudLogTenantWhitelistResponse Response Object
type ShowCloudLogTenantWhitelistResponse struct {

	// 租户Id
	DomainId *string `json:"domain_id,omitempty"`

	// 数仓区域
	DwRegion *string `json:"dw_region,omitempty"`

	// 是否开启
	Enable *bool `json:"enable,omitempty"`

	// 项目Id
	ProjectId      *string `json:"project_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCloudLogTenantWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudLogTenantWhitelistResponse struct{}"
	}

	return strings.Join([]string{"ShowCloudLogTenantWhitelistResponse", string(data)}, " ")
}
