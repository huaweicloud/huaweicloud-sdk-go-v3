package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HourPackage 小时包套餐详情。
type HourPackage struct {

	// 资源所属云服务类型编码。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 小时包的资源规格编码。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 小时包对应的按需桌面的资源规格编码。
	DesktopResourceSpecCode *string `json:"desktop_resource_spec_code,omitempty"`

	Descriptions *ResourcePackageDescription `json:"descriptions,omitempty"`

	// 套餐可使用时长，单位：小时。
	PackageDuration *int32 `json:"package_duration,omitempty"`

	// 该产品套餐支持的专有域id（domainId）。
	DomainIds *[]string `json:"domain_ids,omitempty"`

	// 产品状态，normal：正常、sellout：售空、abandon：下线。
	Status *string `json:"status,omitempty"`
}

func (o HourPackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HourPackage struct{}"
	}

	return strings.Join([]string{"HourPackage", string(data)}, " ")
}
