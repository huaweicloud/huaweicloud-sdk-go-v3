package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VolumeTypeInfo struct {

	// 资源规格编码。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 磁盘类型： - SSD：超高IO磁盘 - SAS：高IO磁盘
	VolumeType *string `json:"volume_type,omitempty"`

	// 磁盘产品类型： - workspace
	VolumeProductType *string `json:"volume_product_type,omitempty"`

	// 资源类型: -hws.resource.type.workspace.volume： 云办公桌面磁盘 -hws.resource.type.workspace.desktop： 云办公桌面 -hws.resource.type.workspace.appstream： 云应用 -hws.resource.type.workspace.appstreamsession： 云应用多会话
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源所属云服务类型编码： - hws.service.type.vdi - hws.service.type.marketplace
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 磁盘中英文名称。
	Name *[]map[string]string `json:"name,omitempty"`

	VolumeTypeExtraSpecs *VolumeTypeExtraSpecs `json:"volume_type_extra_specs,omitempty"`
}

func (o VolumeTypeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeTypeInfo struct{}"
	}

	return strings.Join([]string{"VolumeTypeInfo", string(data)}, " ")
}
