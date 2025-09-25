package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageScanPolicyRespInfoRegistryList 镜像仓库信息
type ImageScanPolicyRespInfoRegistryList struct {

	// 是否已经选择
	IsSelected *bool `json:"is_selected,omitempty"`

	// 镜像仓库名称
	RegistryName *string `json:"registry_name,omitempty"`

	// 镜像仓库id
	RegistryId *string `json:"registry_id,omitempty"`

	// 镜像仓库类型 | SwrPrivate:swr私有 SwrShared:swr共享 SwrEnterprise:swr企业 Harbor:harbor仓库 Jfrog:jfrog仓库 Other:其他仓库
	RegistryType *string `json:"registry_type,omitempty"`
}

func (o ImageScanPolicyRespInfoRegistryList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageScanPolicyRespInfoRegistryList struct{}"
	}

	return strings.Join([]string{"ImageScanPolicyRespInfoRegistryList", string(data)}, " ")
}
