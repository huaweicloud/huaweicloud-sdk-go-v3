package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubPicLayoutInfo struct {

	// 多画面信息。
	Id *int32 `json:"id,omitempty"`

	// 子画面从左到右的索引。
	Left *int32 `json:"left,omitempty"`

	// 子画面从上到下的索引。
	Top *int32 `json:"top,omitempty"`

	// 子画面横向尺寸。
	XSize *int32 `json:"xSize,omitempty"`

	// 子画面横向尺寸。
	YSize *int32 `json:"ySize,omitempty"`
}

func (o SubPicLayoutInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubPicLayoutInfo struct{}"
	}

	return strings.Join([]string{"SubPicLayoutInfo", string(data)}, " ")
}
