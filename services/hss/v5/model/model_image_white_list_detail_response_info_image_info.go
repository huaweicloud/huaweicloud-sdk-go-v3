package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageWhiteListDetailResponseInfoImageInfo struct {

	// **参数解释**： 仓库镜像ID **取值范围**： 取值1-9223372036854775807
	Id *int64 `json:"id,omitempty"`

	// **参数解释**： 本地镜像ID **取值范围**： 字符长度1-64位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**： 镜像名称 **取值范围**： 字符长度1-256位
	ImageName *string `json:"image_name,omitempty"`
}

func (o ImageWhiteListDetailResponseInfoImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWhiteListDetailResponseInfoImageInfo struct{}"
	}

	return strings.Join([]string{"ImageWhiteListDetailResponseInfoImageInfo", string(data)}, " ")
}
