package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageDetectionReq 图像内容审核请求体
type ImageDetectionReq struct {

	// 事件类型。 可选值如下： head_image：头像 album：相册 dynamic：动态 article：帖子 comment：评论 room_cover：房间封面 group_message：群聊图片 message：私聊图片 product：商品图片
	EventType string `json:"event_type"`

	// 检测场景。可添加的检测场景如下： - terrorism：暴恐元素的检测。 - porn：涉黄元素的检测。 - image_text：广告图文的检测。 - 可通过配置上述场景，来完对应场景元素的检测。 > 每个检测场景的检测次数会分类统计。
	Categories []string `json:"categories"`

	ImageTextConfig *ImgTextConfig `json:"image_text_config,omitempty"`

	// 图片url, 与image二选一，目前支持： - 公网HTTP/HTTPS URL
	Url *string `json:"url,omitempty"`

	// 与url二选一，图片文件Base64编码字符串，要求base64编码后大小不超过10M，支持JPG/PNG/JPEG/WEBP/GIF/TIFF/TIF/HEIF等格式。
	Image *string `json:"image,omitempty"`
}

func (o ImageDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionReq struct{}"
	}

	return strings.Join([]string{"ImageDetectionReq", string(data)}, " ")
}
