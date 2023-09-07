package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowIncrementComponentsDetailRequest Request Object
type ShowIncrementComponentsDetailRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowIncrementComponentsDetailRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowIncrementComponentsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIncrementComponentsDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowIncrementComponentsDetailRequest", string(data)}, " ")
}

type ShowIncrementComponentsDetailRequestXLanguage struct {
	value string
}

type ShowIncrementComponentsDetailRequestXLanguageEnum struct {
	EN_US ShowIncrementComponentsDetailRequestXLanguage
	ZH_CN ShowIncrementComponentsDetailRequestXLanguage
}

func GetShowIncrementComponentsDetailRequestXLanguageEnum() ShowIncrementComponentsDetailRequestXLanguageEnum {
	return ShowIncrementComponentsDetailRequestXLanguageEnum{
		EN_US: ShowIncrementComponentsDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowIncrementComponentsDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowIncrementComponentsDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowIncrementComponentsDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowIncrementComponentsDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
