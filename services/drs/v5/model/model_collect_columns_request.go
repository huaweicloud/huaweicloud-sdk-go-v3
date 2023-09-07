package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectColumnsRequest Request Object
type CollectColumnsRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *CollectColumnsRequestXLanguage `json:"X-Language,omitempty"`

	Body *QueryColumnReq `json:"body,omitempty"`
}

func (o CollectColumnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectColumnsRequest struct{}"
	}

	return strings.Join([]string{"CollectColumnsRequest", string(data)}, " ")
}

type CollectColumnsRequestXLanguage struct {
	value string
}

type CollectColumnsRequestXLanguageEnum struct {
	EN_US CollectColumnsRequestXLanguage
	ZH_CN CollectColumnsRequestXLanguage
}

func GetCollectColumnsRequestXLanguageEnum() CollectColumnsRequestXLanguageEnum {
	return CollectColumnsRequestXLanguageEnum{
		EN_US: CollectColumnsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CollectColumnsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CollectColumnsRequestXLanguage) Value() string {
	return c.value
}

func (c CollectColumnsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectColumnsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
