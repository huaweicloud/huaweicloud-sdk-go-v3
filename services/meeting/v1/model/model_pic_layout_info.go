package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PicLayoutInfo struct {
	// 多画面布局的宽度

	X *int32 `json:"x,omitempty"`
	// 多画面布局的高度

	Y *int32 `json:"y,omitempty"`
	// 子画面布局具体列表

	SubPicLayoutInfoList *[]SubPicLayoutInfo `json:"subPicLayoutInfoList,omitempty"`
}

func (o PicLayoutInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PicLayoutInfo struct{}"
	}

	return strings.Join([]string{"PicLayoutInfo", string(data)}, " ")
}
