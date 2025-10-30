package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageData struct {

	// **参数解释** 备份注册镜像ID **取值范围** 字符长度0-256位
	ImageId *string `json:"image_id,omitempty"`
}

func (o ImageData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageData struct{}"
	}

	return strings.Join([]string{"ImageData", string(data)}, " ")
}
