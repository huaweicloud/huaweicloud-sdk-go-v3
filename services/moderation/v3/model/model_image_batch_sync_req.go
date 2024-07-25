package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageBatchSyncReq 图像审核批量同步接口请求参数Body体
type ImageBatchSyncReq struct {

	// 事件类型。可选值如下： - head_image：头像 - album：相册 - dynamic：动态 - article：帖子 - comment：评论 - room_cover：房间封面 - group_message：群聊图片 - message：私聊图片 - product：商品图片
	EventType *string `json:"event_type,omitempty"`

	// 检测场景。可添加的检测场景如下： - terrorism：暴恐元素的检测。 - porn：涉黄元素的检测。 - image_text：广告图文的检测。 可通过配置上述场景，来完对应场景元素的检测。每个检测场景的检测次数会分类统计。
	Categories *[]string `json:"categories,omitempty"`

	ImageTextConfig *ImageBatchSyncReqImageTextConfig `json:"image_text_config,omitempty"`

	// 图片url列表。
	Urls []ImageBatchSyncReqUrls `json:"urls"`

	// 指定图片中文字语种类型。 - zh: 中文 - en: 英文 默认值为zh，中国站仅支持zh。
	Language *string `json:"language,omitempty"`

	// 用户在控制台界面创建的自定义审核策略名称。 - 如果请求参数中包含biz_type则优先使用biztype。 - 如果未传biz_type则event_type和categories为必传参数。
	BizType *string `json:"biz_type,omitempty"`
}

func (o ImageBatchSyncReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageBatchSyncReq struct{}"
	}

	return strings.Join([]string{"ImageBatchSyncReq", string(data)}, " ")
}
