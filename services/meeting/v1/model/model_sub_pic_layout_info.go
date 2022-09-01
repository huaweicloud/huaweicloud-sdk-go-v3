package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubPicLayoutInfo struct {

	// 画面索引号
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 画面从左到右的坐标点
	Left *int32 `json:"left,omitempty" xml:"left"`

	// 画面从上到下的坐标点
	Top *int32 `json:"top,omitempty" xml:"top"`

	// 小画面的宽度
	XSize *int32 `json:"xSize,omitempty" xml:"xSize"`

	// 小画面的高度
	YSize *int32 `json:"ySize,omitempty" xml:"ySize"`
}

func (o SubPicLayoutInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubPicLayoutInfo struct{}"
	}

	return strings.Join([]string{"SubPicLayoutInfo", string(data)}, " ")
}
