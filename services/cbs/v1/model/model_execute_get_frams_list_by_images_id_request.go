package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteGetFramsListByImagesIdRequest struct {

	// 图片id
	ImageId string `json:"image_id"`

	// 偏移量，默认为零
	Offset *int32 `json:"offset,omitempty"`

	// 分页限制数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ExecuteGetFramsListByImagesIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetFramsListByImagesIdRequest struct{}"
	}

	return strings.Join([]string{"ExecuteGetFramsListByImagesIdRequest", string(data)}, " ")
}
