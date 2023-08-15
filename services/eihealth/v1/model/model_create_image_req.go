package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageReq 创建镜像请求体
type CreateImageReq struct {

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 镜像名称
	Name string `json:"name"`

	// 镜像版本
	Tag string `json:"tag"`

	Type *ImageType `json:"type,omitempty"`

	ChipType *ImageChipType `json:"chip_type,omitempty"`
}

func (o CreateImageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageReq struct{}"
	}

	return strings.Join([]string{"CreateImageReq", string(data)}, " ")
}
