package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGetFramsListByImagesIdResponse Response Object
type ExecuteGetFramsListByImagesIdResponse struct {

	// 播报框
	ImageFrames *[]ImageFrame `json:"image_frames,omitempty"`

	// 播报框总数
	Total *int32 `json:"total,omitempty"`

	// 起始偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 本次查询数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExecuteGetFramsListByImagesIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetFramsListByImagesIdResponse struct{}"
	}

	return strings.Join([]string{"ExecuteGetFramsListByImagesIdResponse", string(data)}, " ")
}
