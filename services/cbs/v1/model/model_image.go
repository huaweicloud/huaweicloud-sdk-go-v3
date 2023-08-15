package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Image struct {

	// 创建时间
	CreateTime string `json:"create_time"`

	// 更新时间
	UpdateTime string `json:"update_time"`

	Id string `json:"id"`

	ImageFrame *ImageFrame `json:"image_frame,omitempty"`

	// 图片的obs路径
	ImageUrl string `json:"image_url"`

	Name string `json:"name"`

	// 0：背景图 1：图标 2：系统背景
	Type int32 `json:"type"`

	// 横竖屏模式。0：横版1：竖版
	ResolutionType *int32 `json:"resolution_type,omitempty"`
}

func (o Image) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Image struct{}"
	}

	return strings.Join([]string{"Image", string(data)}, " ")
}
