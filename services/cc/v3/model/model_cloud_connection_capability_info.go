package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudConnectionCapabilityInfo struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 是否支持云连接能力。
	IsSupport bool `json:"is_support"`

	Bandwidth *CloudConnectionDomainBandwidthValue `json:"bandwidth"`

	// 租户支持的区域列表。
	SupportRegions []string `json:"support_regions"`

	// 租户支持的Site列表。
	SupportSites []string `json:"support_sites"`

	ResourceType *CloudConnectionCapabilityKeyEnum `json:"resource_type,omitempty"`
}

func (o CloudConnectionCapabilityInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnectionCapabilityInfo struct{}"
	}

	return strings.Join([]string{"CloudConnectionCapabilityInfo", string(data)}, " ")
}
