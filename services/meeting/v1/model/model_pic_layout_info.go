package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PicLayoutInfo struct {

	// 横向小格子数。
	X *int32 `json:"x,omitempty"`

	// 纵向小格子数。
	Y *int32 `json:"y,omitempty"`

	// 多画面信息。
	SubPicLayoutInfoList *[]SubPicLayoutInfo `json:"subPicLayoutInfoList,omitempty"`
}

func (o PicLayoutInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PicLayoutInfo struct{}"
	}

	return strings.Join([]string{"PicLayoutInfo", string(data)}, " ")
}
