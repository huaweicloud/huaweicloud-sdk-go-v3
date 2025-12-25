package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageDetectionReq 图像内容审核请求体
type ImageDetectionReq struct {

	// 事件类型。 可选值如下： head_image：头像 album：相册 dynamic：动态 article：帖子 comment：评论 room_cover：房间封面 group_message：群聊图片 message：私聊图片 product：商品图片
	EventType *string `json:"event_type,omitempty"`

	// 检测场景。可添加的检测场景如下： - terrorism：暴恐元素的检测。 - porn：涉黄元素的检测。 - image_text：广告图文的检测。 - politics: 涉政人物的检测。 - 可通过配置上述场景，来完对应场景元素的检测。 > 每个检测场景的检测次数会分类统计。
	Categories *[]string `json:"categories,omitempty"`

	ImageTextConfig *ImgTextConfig `json:"image_text_config,omitempty"`

	// 图片url, 与image二选一，目前支持： - 公网HTTP/HTTPS URL
	Url *string `json:"url,omitempty"`

	// 与url二选一，图片文件Base64编码字符串，要求base64编码后大小不超过10M，支持JPG/PNG/JPEG/WEBP/GIF/TIFF/TIF/HEIF等格式。
	Image *string `json:"image,omitempty"`

	// 自定义审核策略名称，可在控制台配置;如果请求参数中传了biz_type则优先使用biz_type,如果用户没传biz_type则event_type和categories必须传。
	BizType *string `json:"biz_type,omitempty"`

	// 可指定图片中文字语种类型
	Language *ImageDetectionReqLanguage `json:"language,omitempty"`
}

func (o ImageDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionReq struct{}"
	}

	return strings.Join([]string{"ImageDetectionReq", string(data)}, " ")
}

type ImageDetectionReqLanguage struct {
	value string
}

type ImageDetectionReqLanguageEnum struct {
	ZH ImageDetectionReqLanguage
}

func GetImageDetectionReqLanguageEnum() ImageDetectionReqLanguageEnum {
	return ImageDetectionReqLanguageEnum{
		ZH: ImageDetectionReqLanguage{
			value: "zh",
		},
	}
}

func (c ImageDetectionReqLanguage) Value() string {
	return c.value
}

func (c ImageDetectionReqLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageDetectionReqLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
