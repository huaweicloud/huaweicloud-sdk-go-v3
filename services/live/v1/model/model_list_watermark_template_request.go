package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWatermarkTemplateRequest Request Object
type ListWatermarkTemplateRequest struct {

	// 水印模板名称
	Name *string `json:"name,omitempty"`

	// 区分媒体直播还是云直播，默认云直播，默认查询cloud_live
	Scene *ListWatermarkTemplateRequestScene `json:"scene,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，取值范围[1,100]，默认值10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListWatermarkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListWatermarkTemplateRequest", string(data)}, " ")
}

type ListWatermarkTemplateRequestScene struct {
	value string
}

type ListWatermarkTemplateRequestSceneEnum struct {
	CLOUD_LIVE ListWatermarkTemplateRequestScene
	MEDIA_LIVE ListWatermarkTemplateRequestScene
}

func GetListWatermarkTemplateRequestSceneEnum() ListWatermarkTemplateRequestSceneEnum {
	return ListWatermarkTemplateRequestSceneEnum{
		CLOUD_LIVE: ListWatermarkTemplateRequestScene{
			value: "cloud_live",
		},
		MEDIA_LIVE: ListWatermarkTemplateRequestScene{
			value: "media_live",
		},
	}
}

func (c ListWatermarkTemplateRequestScene) Value() string {
	return c.value
}

func (c ListWatermarkTemplateRequestScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWatermarkTemplateRequestScene) UnmarshalJSON(b []byte) error {
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
