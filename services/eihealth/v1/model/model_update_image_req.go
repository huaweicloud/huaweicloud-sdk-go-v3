package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageReq 更新镜像请求体
type UpdateImageReq struct {
	Type *ImageType `json:"type,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	ChipType *ImageChipType `json:"chip_type,omitempty"`
}

func (o UpdateImageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageReq struct{}"
	}

	return strings.Join([]string{"UpdateImageReq", string(data)}, " ")
}
