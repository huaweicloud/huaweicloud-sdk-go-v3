package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchBoxDetail struct {

	// 区域中x坐标的最小值，单位：像素。
	X *int32 `json:"x,omitempty"`

	// 区域中y坐标的最小值，单位：像素。
	Y *int32 `json:"y,omitempty"`

	// 区域的宽度，单位：像素。
	Width *int32 `json:"width,omitempty"`

	// 区域的高度，单位：像素。
	Height *int32 `json:"height,omitempty"`
}

func (o SearchBoxDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchBoxDetail struct{}"
	}

	return strings.Join([]string{"SearchBoxDetail", string(data)}, " ")
}
