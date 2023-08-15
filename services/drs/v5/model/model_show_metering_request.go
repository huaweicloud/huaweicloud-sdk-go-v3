package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMeteringRequest Request Object
type ShowMeteringRequest struct {

	// 请求语言类型。
	XLanguage *ShowMeteringRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ShowMeteringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMeteringRequest struct{}"
	}

	return strings.Join([]string{"ShowMeteringRequest", string(data)}, " ")
}

type ShowMeteringRequestXLanguage struct {
	value string
}

type ShowMeteringRequestXLanguageEnum struct {
	EN_US ShowMeteringRequestXLanguage
	ZH_CN ShowMeteringRequestXLanguage
}

func GetShowMeteringRequestXLanguageEnum() ShowMeteringRequestXLanguageEnum {
	return ShowMeteringRequestXLanguageEnum{
		EN_US: ShowMeteringRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowMeteringRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowMeteringRequestXLanguage) Value() string {
	return c.value
}

func (c ShowMeteringRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMeteringRequestXLanguage) UnmarshalJSON(b []byte) error {
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
