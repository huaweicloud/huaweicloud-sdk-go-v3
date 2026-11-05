package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageUploadItem struct {

	// 图片名称（后缀必须为jpg/jpeg/png/bmp）
	ImageName string `json:"image_name"`

	// 图片大小（字节，不超过max_image_size默认5MB）
	ImageSize int64 `json:"image_size"`

	// 上传内容类型（作为Content-Type请求头纳入OBS预签名，客户端PUT上传时须用相同值）
	UploadContentType string `json:"upload_content_type"`
}

func (o ImageUploadItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageUploadItem struct{}"
	}

	return strings.Join([]string{"ImageUploadItem", string(data)}, " ")
}
