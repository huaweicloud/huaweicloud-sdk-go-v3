package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateManualImageScanTaskReqInfoImageInfo struct {

	// id
	Id *int64 `json:"id,omitempty"`

	// image id
	ImageId *string `json:"image_id,omitempty"`

	// 镜像digest
	ImageDigest *string `json:"image_digest,omitempty"`

	// 镜像仓库ID
	RegistryId *string `json:"registry_id,omitempty"`

	// 镜像仓库名称
	RegistryName *string `json:"registry_name,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 企业实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 下载企业镜像URL
	InstanceUrl *string `json:"instance_url,omitempty"`

	// **参数解释**: 镜像仓库类型 **约束限制**: 不涉及 **取值范围**: - SwrPrivate：swr私有。 - SwrShared：swr共享。 - SwrEnterprise：swr企业。 - Harbor：harbor仓库。 - Jfrog：jfrog仓库。 - Other：其他仓库。  **默认取值**: 不涉及
	RegistryType *string `json:"registry_type,omitempty"`
}

func (o CreateManualImageScanTaskReqInfoImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualImageScanTaskReqInfoImageInfo struct{}"
	}

	return strings.Join([]string{"CreateManualImageScanTaskReqInfoImageInfo", string(data)}, " ")
}
