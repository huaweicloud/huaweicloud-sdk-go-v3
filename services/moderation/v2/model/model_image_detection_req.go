package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type ImageDetectionReq struct {

	// 与image二选一  图片的URL路径，目前支持：  - 公网HTTP/HTTPS URL  - 华为云OBS提供的URL，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详请参见[配置OBS访问权限](https://support.huaweicloud.com/api-moderation/moderation_03_0020.html)。   > - 接口响应时间依赖图片的下载时间，如果图片下载时间过长，会返回接口调用失败。请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。
	Url *string `json:"url,omitempty"`

	// 与url二选一  图片文件Base64编码字符串。要求base64编码后大小不超过10M。  政治人物检测人脸部分不小于50*50像素。  支持JPG/PNG/BMP格式。
	Image *string `json:"image,omitempty"`

	// 图像审核规则名称，默认使用default规则。  > 审核规则的创建和使用请参见[配置审核规则](https://support.huaweicloud.com/api-moderation/moderation_03_0063.html)
	ModerationRule *string `json:"moderation_rule,omitempty"`

	// 检测场景:  - politics：是否涉及政治人物的检测。  - terrorism：是否包含涉政暴恐元素的检测。  - porn：是否包含涉黄内容元素的检测。  - ad：是否包含广告的检测（公测特性）。  - all：包含politics、terrorism和porn三种场景的检测。  可通过配置上述场景，来完对应场景元素的检测。  为空或无此参数表示politics和terrorism都检测，但不包含porn场景。  > 每个检测场景的检测次数会分类统计。
	Categories *[]ImageDetectionReqCategories `json:"categories,omitempty"`

	// 图文审核检测场景。当categories包含ad时，该参数生效。 当前支持的场景有系统场景和用户自定义场景： - 系统场景为：   - qr_code：二维码   - politics：涉政   - porn：涉黄   - ad：广告   - abuse：辱骂   - contraband：违禁品 - 用户自定义场景为：自定义黑名单词库。   > 自定义词库的创建和使用请参见[配置自定义词库](https://support.huaweicloud.com/api-moderation/moderation_03_0020.html)。
	AdCategories *[]string `json:"ad_categories,omitempty"`

	// - 结果过滤门限，只有置信度不低于此门限的结果才会呈现在detail的列表中，取值范围 0-1，当未设置此值时各个检测场景会使用各自的默认值。  - politics检测场景的默认值为0.95。  - terrorism检测场景的默认值为0。  - ad检测场景的默认值为0。  - 无特殊需求直接不传此参数或像示例中一样值设为空字符串即可。  > - 如果检测场景中的最高置信度也未达到threshold，则结果列表为空；反之threshold过小，则会使结果列表中内容过多。 > - threshold参数不支持porn场景筛选。 > -  threshold参数不会影响响应中的suggestion。
	Threshold *float32 `json:"threshold,omitempty"`

	// 是否返回ocr识别结果，默认为false。
	ShowOcrText *bool `json:"show_ocr_text,omitempty"`
}

func (o ImageDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionReq struct{}"
	}

	return strings.Join([]string{"ImageDetectionReq", string(data)}, " ")
}

type ImageDetectionReqCategories struct {
	value string
}

type ImageDetectionReqCategoriesEnum struct {
	POLITICS  ImageDetectionReqCategories
	TERRORISM ImageDetectionReqCategories
	PORN      ImageDetectionReqCategories
	AD        ImageDetectionReqCategories
	ALL       ImageDetectionReqCategories
}

func GetImageDetectionReqCategoriesEnum() ImageDetectionReqCategoriesEnum {
	return ImageDetectionReqCategoriesEnum{
		POLITICS: ImageDetectionReqCategories{
			value: "politics",
		},
		TERRORISM: ImageDetectionReqCategories{
			value: "terrorism",
		},
		PORN: ImageDetectionReqCategories{
			value: "porn",
		},
		AD: ImageDetectionReqCategories{
			value: "ad",
		},
		ALL: ImageDetectionReqCategories{
			value: "all",
		},
	}
}

func (c ImageDetectionReqCategories) Value() string {
	return c.value
}

func (c ImageDetectionReqCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageDetectionReqCategories) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
