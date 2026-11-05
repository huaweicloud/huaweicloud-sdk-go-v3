package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AtlasImageItem struct {

	// 图片ID
	ImageId *string `json:"image_id,omitempty"`

	// 图片名称
	Name *string `json:"name,omitempty"`

	// 图片查看的预签名下载地址（GET）
	Url *string `json:"url,omitempty"`

	// 图片创建时间（毫秒时间戳）
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o AtlasImageItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtlasImageItem struct{}"
	}

	return strings.Join([]string{"AtlasImageItem", string(data)}, " ")
}
