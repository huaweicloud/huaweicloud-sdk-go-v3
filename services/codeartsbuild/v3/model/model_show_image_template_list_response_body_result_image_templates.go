package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageTemplateListResponseBodyResultImageTemplates 镜像模版信息
type ShowImageTemplateListResponseBodyResultImageTemplates struct {

	// 镜像模版ID
	ImageId *string `json:"image_id,omitempty"`

	// swr镜像组织
	Organization *string `json:"organization,omitempty"`

	// 镜像名
	ImageName *string `json:"image_name,omitempty"`

	// 镜像label
	ImageLabel *string `json:"image_label,omitempty"`

	// 操作系统
	BaseImage *string `json:"base_image,omitempty"`

	// 镜像标题
	Title *string `json:"title,omitempty"`

	// 镜像描述
	Description *string `json:"description,omitempty"`

	// 镜像创建日期
	CreateTime *string `json:"create_time,omitempty"`
}

func (o ShowImageTemplateListResponseBodyResultImageTemplates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageTemplateListResponseBodyResultImageTemplates struct{}"
	}

	return strings.Join([]string{"ShowImageTemplateListResponseBodyResultImageTemplates", string(data)}, " ")
}
