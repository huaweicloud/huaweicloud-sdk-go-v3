package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageBasicImageInfo struct {

	// 基础镜像的系统名称
	OsName *string `json:"os_name,omitempty"`

	// 基础镜像的系统版本
	OsVersion *string `json:"os_version,omitempty"`

	// 基础镜像的层digest
	LayerDigest *string `json:"layer_digest,omitempty"`
}

func (o ImageBasicImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageBasicImageInfo struct{}"
	}

	return strings.Join([]string{"ImageBasicImageInfo", string(data)}, " ")
}
