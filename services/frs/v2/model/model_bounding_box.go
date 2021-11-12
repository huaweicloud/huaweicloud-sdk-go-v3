package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BoundingBox struct {
	// 矩形框宽度。

	Width int32 `json:"width"`
	// 矩形框左上角纵坐标。

	TopLeftY int32 `json:"top_left_y"`
	// 矩形框左上角横坐标。

	TopLeftX int32 `json:"top_left_x"`
	// 矩形框高度。

	Height int32 `json:"height"`
}

func (o BoundingBox) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BoundingBox struct{}"
	}

	return strings.Join([]string{"BoundingBox", string(data)}, " ")
}
