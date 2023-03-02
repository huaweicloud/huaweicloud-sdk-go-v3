package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto crerate Body Object
type ImageWisedesignCombineReq struct {

	// 画布的宽度
	CanvasW int32 `json:"canvas_w"`

	// 画布的高度
	CanvasH int32 `json:"canvas_h"`

	// 图片合成的层级详细信息
	Layers []ImageWisedesignCombineBody `json:"layers"`
}

func (o ImageWisedesignCombineReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWisedesignCombineReq struct{}"
	}

	return strings.Join([]string{"ImageWisedesignCombineReq", string(data)}, " ")
}
