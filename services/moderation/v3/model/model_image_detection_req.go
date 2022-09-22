package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 图像内容审核请求体
type ImageDetectionReq struct {

	// 事件类型。 可选值如下： head_image：头像 album：相册 dynamic：动态 article：帖子 comment：评论 room_cover：房间封面 group_message：群聊图片 message：私聊图片 product：商品图片
	EventType string `json:"event_type"`

	// 检测场景，可添加的检测场景如下，华为云当前支持全场景的，以下仅展示部门检测能力，如果有其它检测场景的需求，请在开通服务时咨询华为云工程师：  terrorism：涉政暴恐内容的检测。  porn：鉴黄内容的检测。  politics：政治敏感人物内容的检测 image_text：图文违规内容的检测。（检测图片中出现的广告、色情、暴恐、涉政的文字违规内容以及二维码内容）  备注：资料上线需要删除politics
	Categories []string `json:"categories"`

	ImageTextConfig *ImgTextConfig `json:"image_text_config,omitempty"`

	// 图片url, 与image二选一。 图片的URL路径，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详情请见配置OBS访问权限。 ​ 说明： 接口响应时间依赖图片的下载时间，如果图片下载时间过长，会返回接口调用失败。请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。 图片url, 与image二选一。 图片的URL路径，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详情请见配置OBS访问权限。  说明： 接口响应时间依赖图片的下载时间，如果图片下载时间过长，会返回接口调用失败。请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。
	Url *string `json:"url,omitempty"`

	// 与url二选一。 ​ 图片文件Base64编码字符串。要求base64编码后大小不超过10M。 ​ 支持JPG/PNG/BMP/WEBP等格式。 与url二选一。  图片文件Base64编码字符串。要求base64编码后大小不超过10M。  支持JPG/PNG/BMP/WEBP等格式。
	Image *string `json:"image,omitempty"`
}

func (o ImageDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionReq struct{}"
	}

	return strings.Join([]string{"ImageDetectionReq", string(data)}, " ")
}
