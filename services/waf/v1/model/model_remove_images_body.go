package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveImagesBody struct {

	// 需移除的图片ID列表
	ImageIds []string `json:"image_ids"`
}

func (o RemoveImagesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveImagesBody struct{}"
	}

	return strings.Join([]string{"RemoveImagesBody", string(data)}, " ")
}
