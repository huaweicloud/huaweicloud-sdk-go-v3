package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageScanTaskInfoFailedImages 失败镜像信息
type ImageScanTaskInfoFailedImages struct {

	// **参数解释**： 失败镜像ID **取值范围**： 最小值0，最大值9223372036854775807
	Id *int64 `json:"id,omitempty"`

	// **参数解释**： 镜像仓库ID **取值范围**： 字符长度0-128位
	RegistryId *string `json:"registry_id,omitempty"`

	// **参数解释**： 镜像仓库名称 **取值范围**： 字符长度0-128位
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**： 镜像名称 **取值范围**： 字符长度0-128位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**： 镜像版本 **取值范围**： 字符长度0-128位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**： 命名空间 **取值范围**： 字符长度0-128位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 镜像仓库类型 **取值范围**： - SwrPrivate：swr私有。 - SwrShared：swr共享。 - SwrEnterprise：swr企业。 - Harbor：harbor仓库。 - Jfrog：jfrog仓库。 - Other：其他仓库。
	RegistryType *string `json:"registry_type,omitempty"`

	// **参数解释**： 失败原因 **取值范围**： 字符长度0-128位
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o ImageScanTaskInfoFailedImages) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageScanTaskInfoFailedImages struct{}"
	}

	return strings.Join([]string{"ImageScanTaskInfoFailedImages", string(data)}, " ")
}
