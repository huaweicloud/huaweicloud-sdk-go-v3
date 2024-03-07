package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImagesRequest Request Object
type ListImagesRequest struct {

	// 产品镜像的操作系统类型，如Windows。
	OsType *string `json:"os_type,omitempty"`

	// 镜像类型。 -gold  公共镜像 -private  私有镜像
	ImageType *string `json:"image_type,omitempty"`

	// 镜像系统类型,如Windows。
	Platform *string `json:"platform,omitempty"`

	// 镜像架构：x86。
	Architecture *string `json:"architecture,omitempty"`

	// 套餐系列
	PackageType *string `json:"package_type,omitempty"`

	// 镜像Id
	ImageId *string `json:"image_id,omitempty"`

	// 每页数量，范围0-100，默认100。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量,默认0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesRequest struct{}"
	}

	return strings.Join([]string{"ListImagesRequest", string(data)}, " ")
}
