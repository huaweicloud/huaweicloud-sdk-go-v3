package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImagesView 镜像详情
type ListImagesView struct {

	// 镜像所属租户
	DomainId *string `json:"domain_id,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 镜像创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 镜像大小，单位byte
	ImageSize *int64 `json:"image_size,omitempty"`

	// project_id（当image_type为private时，才会返回此字段)
	ProjectId *string `json:"project_id,omitempty"`

	// 镜像ID
	ImageId *string `json:"image_id,omitempty"`

	// 镜像AOSP版本
	ImageVersion *string `json:"image_version,omitempty"`

	// 镜像类型 公共镜像：public 私有镜像：private 共享镜像：share
	ImageType *string `json:"image_type,omitempty"`

	// 镜像状态。 - 0：CREATING 创建中 - 1：PRODUCTION 生产态，可使用 - 2：CREATE_FAILED 创建失败
	Status *int32 `json:"status,omitempty"`

	// 共享镜像账号的projectId（当image_type为share时，才会返回此字段)
	SrcProjectId *string `json:"src_project_id,omitempty"`
}

func (o ListImagesView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesView struct{}"
	}

	return strings.Join([]string{"ListImagesView", string(data)}, " ")
}
