package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 镜像元素
type ImageData struct {

	// 镜像ID
	ImageId *string `json:"image_id,omitempty" xml:"image_id"`
}

func (o ImageData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageData struct{}"
	}

	return strings.Join([]string{"ImageData", string(data)}, " ")
}
