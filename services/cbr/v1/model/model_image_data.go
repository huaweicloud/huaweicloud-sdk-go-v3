package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 镜像元素
type ImageData struct {
	//

	ImageId *string `json:"image_id,omitempty"`
}

func (o ImageData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageData struct{}"
	}

	return strings.Join([]string{"ImageData", string(data)}, " ")
}
