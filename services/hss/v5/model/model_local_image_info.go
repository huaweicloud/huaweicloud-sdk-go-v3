package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LocalImageInfo struct {

	// 镜像id
	ImageId *string `json:"image_id,omitempty"`

	// 镜像摘要
	ImageDigest *string `json:"image_digest,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`
}

func (o LocalImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalImageInfo struct{}"
	}

	return strings.Join([]string{"LocalImageInfo", string(data)}, " ")
}
