package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageInfo 镜像信息
type ImageInfo struct {

	// 源项目名
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 镜像名
	ImageName *string `json:"image_name,omitempty"`

	// 镜像tag名
	ImageTag *string `json:"image_tag,omitempty"`

	Profile *Profile `json:"profile,omitempty"`
}

func (o ImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageInfo struct{}"
	}

	return strings.Join([]string{"ImageInfo", string(data)}, " ")
}
