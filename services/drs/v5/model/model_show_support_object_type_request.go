package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSupportObjectTypeRequest Request Object
type ShowSupportObjectTypeRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowSupportObjectTypeRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowSupportObjectTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupportObjectTypeRequest struct{}"
	}

	return strings.Join([]string{"ShowSupportObjectTypeRequest", string(data)}, " ")
}

type ShowSupportObjectTypeRequestXLanguage struct {
	value string
}

type ShowSupportObjectTypeRequestXLanguageEnum struct {
	EN_US ShowSupportObjectTypeRequestXLanguage
	ZH_CN ShowSupportObjectTypeRequestXLanguage
}

func GetShowSupportObjectTypeRequestXLanguageEnum() ShowSupportObjectTypeRequestXLanguageEnum {
	return ShowSupportObjectTypeRequestXLanguageEnum{
		EN_US: ShowSupportObjectTypeRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowSupportObjectTypeRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowSupportObjectTypeRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSupportObjectTypeRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSupportObjectTypeRequestXLanguage) UnmarshalJSON(b []byte) error {
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
