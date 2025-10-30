package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulRepoImageInfo struct {

	// **参数解释**: 组织名称 **取值范围**: 字符长度0-65535位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像id **取值范围**: 字符长度0-65535位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-65535位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 应用名称 **取值范围**: 字符长度0-65535位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 版本信息 **取值范围**: 字符长度0-65535位
	Version *string `json:"version,omitempty"`

	// **参数解释**: 镜像版本数 **取值范围**: 最小值0，最大值65535
	TagCount *int32 `json:"tag_count,omitempty"`

	// **参数解释**: 镜像类型 **取值范围**: - private_image：私有镜像仓库。 - shared_image：共享镜像仓库。
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: tag版本总数 **取值范围**: 最小值0，最大值65535
	DataList *[]VulRepoImagesTagInfo `json:"data_list,omitempty"`
}

func (o VulRepoImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulRepoImageInfo struct{}"
	}

	return strings.Join([]string{"VulRepoImageInfo", string(data)}, " ")
}
