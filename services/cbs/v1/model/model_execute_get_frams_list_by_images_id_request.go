package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteGetFramsListByImagesIdRequest struct {

	// 图片id
	ImageId string `json:"image_id"`
}

func (o ExecuteGetFramsListByImagesIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetFramsListByImagesIdRequest struct{}"
	}

	return strings.Join([]string{"ExecuteGetFramsListByImagesIdRequest", string(data)}, " ")
}
