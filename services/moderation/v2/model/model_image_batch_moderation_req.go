package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type ImageBatchModerationReq struct {
	// 图片的URL路径，目前支持： - 公网HTTP/HTTPS URL - 华为云OBS提供的URL，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详请参见[配置OBS访问权限](https://support.huaweicloud.com/api-moderation/moderation_03_0020.html)。 > 图片的URL路径列表最多支持10个URL地址。接口响应时间依赖图片的下载时间，如果图片下载时间过长，会返回接口调用失败。请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。

	Urls []string `json:"urls"`
	// 检测场景:  - politics：是否涉及政治人物的检测。  - terrorism：是否包含涉政暴恐元素的检测。  - porn：是否包含涉黄内容元素的检测。  - ad：是否包含广告的检测（公测特性）。  - all：包含politics、terrorism和porn三种场景的检测。  可通过配置上述场景，来完对应场景元素的检测。  为空或无此参数表示politics和terrorism都检测，但不包含porn场景。  > 每个检测场景的检测次数会分类统计。

	Categories *[]ImageBatchModerationReqCategories `json:"categories,omitempty"`
	// - 结果过滤门限，只有置信度不低于此门限的结果才会呈现在detail的列表中，取值范围 0-1，当未设置此值时各个检测场景会使用各自的默认值。  - politics检测场景的默认值为0.95。  - terrorism检测场景的默认值为0。  - ad检测场景的默认值为0。  - 无特殊需求直接不传此参数或像示例中一样值设为空字符串即可。  > - 如果检测场景中的最高置信度也未达到threshold，则结果列表为空；反之threshold过小，则会使结果列表中内容过多。 > - threshold参数不支持porn场景筛选。 > -  threshold参数不会影响响应中的suggestion。

	Threshold *float64 `json:"threshold,omitempty"`
}

func (o ImageBatchModerationReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageBatchModerationReq struct{}"
	}

	return strings.Join([]string{"ImageBatchModerationReq", string(data)}, " ")
}

type ImageBatchModerationReqCategories struct {
	value string
}

type ImageBatchModerationReqCategoriesEnum struct {
	POLITICS  ImageBatchModerationReqCategories
	TERRORISM ImageBatchModerationReqCategories
	PORN      ImageBatchModerationReqCategories
	AD        ImageBatchModerationReqCategories
	ALL       ImageBatchModerationReqCategories
}

func GetImageBatchModerationReqCategoriesEnum() ImageBatchModerationReqCategoriesEnum {
	return ImageBatchModerationReqCategoriesEnum{
		POLITICS: ImageBatchModerationReqCategories{
			value: "politics",
		},
		TERRORISM: ImageBatchModerationReqCategories{
			value: "terrorism",
		},
		PORN: ImageBatchModerationReqCategories{
			value: "porn",
		},
		AD: ImageBatchModerationReqCategories{
			value: "ad",
		},
		ALL: ImageBatchModerationReqCategories{
			value: "all",
		},
	}
}

func (c ImageBatchModerationReqCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ImageBatchModerationReqCategories) UnmarshalJSON(b []byte) error {
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
