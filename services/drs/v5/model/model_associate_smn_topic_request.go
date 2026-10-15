package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssociateSmnTopicRequest Request Object
type AssociateSmnTopicRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *AssociateSmnTopicRequestXLanguage `json:"X-Language,omitempty"`

	Body *AssociateSmnTopicReqV5 `json:"body,omitempty"`
}

func (o AssociateSmnTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSmnTopicRequest struct{}"
	}

	return strings.Join([]string{"AssociateSmnTopicRequest", string(data)}, " ")
}

type AssociateSmnTopicRequestXLanguage struct {
	value string
}

type AssociateSmnTopicRequestXLanguageEnum struct {
	EN_US AssociateSmnTopicRequestXLanguage
	ZH_CN AssociateSmnTopicRequestXLanguage
	ES_US AssociateSmnTopicRequestXLanguage
	PT_BR AssociateSmnTopicRequestXLanguage
}

func GetAssociateSmnTopicRequestXLanguageEnum() AssociateSmnTopicRequestXLanguageEnum {
	return AssociateSmnTopicRequestXLanguageEnum{
		EN_US: AssociateSmnTopicRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: AssociateSmnTopicRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: AssociateSmnTopicRequestXLanguage{
			value: "es-us",
		},
		PT_BR: AssociateSmnTopicRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c AssociateSmnTopicRequestXLanguage) Value() string {
	return c.value
}

func (c AssociateSmnTopicRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssociateSmnTopicRequestXLanguage) UnmarshalJSON(b []byte) error {
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
