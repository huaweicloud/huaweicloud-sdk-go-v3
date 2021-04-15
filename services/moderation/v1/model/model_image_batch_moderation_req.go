package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type ImageBatchModerationReq struct {
	// 图片的URL路径，目前支持： - 公网HTTP/HTTPS URL - 华为云OBS提供的URL，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详请参见[配置OBS访问权限](https://support.huaweicloud.com/api-moderation/moderation_03_0020.html)。 > 说明：图片的URL路径列表最多支持10个URL地址。接口响应时间依赖图片的下载时间，如果图片下载时间过长，会返回接口调用失败。请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。

	Urls []string `json:"urls"`
	// 请参见[表1](https://support.huaweicloud.com/api-moderation/moderation_03_0019.html#moderation_03_0019__zh-cn_topic_0087142444_table2693546819540)中categories参数说明。

	Categories *[]ImageBatchModerationReqCategories `json:"categories,omitempty"`
	// 请参见[表1](https://support.huaweicloud.com/api-moderation/moderation_03_0019.html#moderation_03_0019__zh-cn_topic_0087142444_table2693546819540)中threshold参数说明。

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
