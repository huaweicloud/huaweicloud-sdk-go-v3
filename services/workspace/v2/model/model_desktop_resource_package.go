package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DesktopResourcePackage struct {

	// 云服务类型。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源规格编码。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 规格。
	PackageType *string `json:"package_type,omitempty"`

	// cpu。
	Cpu *string `json:"cpu,omitempty"`

	// 产品架构：arm、x86。
	Architecture *string `json:"architecture,omitempty"`

	// 内存。
	Memory *string `json:"memory,omitempty"`

	// 该产品套餐支持的专有域id（domainId）。
	DomainIds *[]string `json:"domain_ids,omitempty"`

	// 产品名称<语言，各语言对应的产品名>。
	Name map[string]string `json:"name,omitempty"`

	// 产品描述<语言，各语言对应的产品名>。
	Description map[string]string `json:"description,omitempty"`

	// 按需套餐包规格列表。
	ResourcePackages *[]ResourcePackage `json:"resource_packages,omitempty"`
}

func (o DesktopResourcePackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopResourcePackage struct{}"
	}

	return strings.Join([]string{"DesktopResourcePackage", string(data)}, " ")
}
