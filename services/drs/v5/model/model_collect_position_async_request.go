package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectPositionAsyncRequest Request Object
type CollectPositionAsyncRequest struct {

	// 任务ID
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *CollectPositionAsyncRequestXLanguage `json:"X-Language,omitempty"`

	Body *QueryDbPositionReq `json:"body,omitempty"`
}

func (o CollectPositionAsyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectPositionAsyncRequest struct{}"
	}

	return strings.Join([]string{"CollectPositionAsyncRequest", string(data)}, " ")
}

type CollectPositionAsyncRequestXLanguage struct {
	value string
}

type CollectPositionAsyncRequestXLanguageEnum struct {
	EN_US CollectPositionAsyncRequestXLanguage
	ZH_CN CollectPositionAsyncRequestXLanguage
}

func GetCollectPositionAsyncRequestXLanguageEnum() CollectPositionAsyncRequestXLanguageEnum {
	return CollectPositionAsyncRequestXLanguageEnum{
		EN_US: CollectPositionAsyncRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CollectPositionAsyncRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CollectPositionAsyncRequestXLanguage) Value() string {
	return c.value
}

func (c CollectPositionAsyncRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectPositionAsyncRequestXLanguage) UnmarshalJSON(b []byte) error {
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
