package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryAppImageResponseDto 应用镜像列表信息
type QueryAppImageResponseDto struct {

	// 镜像组织
	ImageNamespace *string `json:"image_namespace,omitempty"`

	// 镜像仓库名称
	Name *string `json:"name,omitempty"`

	// 镜像tag
	Tag *string `json:"tag,omitempty"`

	// 镜像摘要
	Digest *string `json:"digest,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`
}

func (o QueryAppImageResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAppImageResponseDto struct{}"
	}

	return strings.Join([]string{"QueryAppImageResponseDto", string(data)}, " ")
}
