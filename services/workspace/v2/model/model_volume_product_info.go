package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeProductInfo 磁盘产品信息
type VolumeProductInfo struct {

	// 产品ID
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 磁盘类型： - SATA: 普通IO磁盘 - SAS：高IO磁盘 - SSD：超高IO磁盘
	VolumeType *string `json:"volume_type,omitempty"`

	// 产品类型：workspace
	VolumeProductType *string `json:"volume_product_type,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 云服务类型
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 该磁盘支持的专有域id（domainId）。
	DomainIds *[]string `json:"domain_ids,omitempty"`

	// 磁盘名称
	Name *[]I18nLanguage `json:"name,omitempty"`

	// 产品状态，normal：正常、sellout：售空。
	Status *string `json:"status,omitempty"`
}

func (o VolumeProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeProductInfo struct{}"
	}

	return strings.Join([]string{"VolumeProductInfo", string(data)}, " ")
}
