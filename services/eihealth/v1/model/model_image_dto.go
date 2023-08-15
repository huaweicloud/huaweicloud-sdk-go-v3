package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageDto struct {

	// 镜像类型
	Type *string `json:"type,omitempty"`

	// 镜像版本
	Version *string `json:"version,omitempty"`
}

func (o ImageDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDto struct{}"
	}

	return strings.Join([]string{"ImageDto", string(data)}, " ")
}
