package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperImageOptionsResponseInfo 网页防篡改可选服务器信息
type WebTamperImageOptionsResponseInfo struct {

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-512位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-512位
	ImageFullName *string `json:"image_full_name,omitempty"`

	// **参数解释**: 镜像id **取值范围**: 字符长度0-64位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 镜像版本列表 **约束限制**: 不涉及 **取值范围**: 最少0条，最多100条 **默认取值**: 不涉及
	ImageVersionList *[]string `json:"image_version_list,omitempty"`

	// **参数解释**: 仓库镜像的组织名称 **取值范围**: 字符长度0-512位
	ImageNamespace *string `json:"image_namespace,omitempty"`

	// **参数解释**: 镜像仓库名称 **取值范围**: 字符长度1-128位
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**： 镜像仓库类型 **取值范围**： - SwrPrivate：swr私有。 - SwrShared：swr共享。 - SwrEnterprise：swr企业。 - Harbor：harbor仓库。 - Jfrog：jfrog仓库。 - Other：其他仓库。
	RegistryType *string `json:"registry_type,omitempty"`
}

func (o WebTamperImageOptionsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperImageOptionsResponseInfo struct{}"
	}

	return strings.Join([]string{"WebTamperImageOptionsResponseInfo", string(data)}, " ")
}
