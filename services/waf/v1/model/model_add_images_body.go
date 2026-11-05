package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddImagesBody struct {

	// 需加入的图片ID列表
	ImageIds []string `json:"image_ids"`
}

func (o AddImagesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImagesBody struct{}"
	}

	return strings.Join([]string{"AddImagesBody", string(data)}, " ")
}
