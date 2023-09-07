package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckDataFilterRequest Request Object
type CheckDataFilterRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *CheckDataFilterRequestXLanguage `json:"X-Language,omitempty"`

	Body *DataProcessReq `json:"body,omitempty"`
}

func (o CheckDataFilterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataFilterRequest struct{}"
	}

	return strings.Join([]string{"CheckDataFilterRequest", string(data)}, " ")
}

type CheckDataFilterRequestXLanguage struct {
	value string
}

type CheckDataFilterRequestXLanguageEnum struct {
	EN_US CheckDataFilterRequestXLanguage
	ZH_CN CheckDataFilterRequestXLanguage
}

func GetCheckDataFilterRequestXLanguageEnum() CheckDataFilterRequestXLanguageEnum {
	return CheckDataFilterRequestXLanguageEnum{
		EN_US: CheckDataFilterRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CheckDataFilterRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CheckDataFilterRequestXLanguage) Value() string {
	return c.value
}

func (c CheckDataFilterRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckDataFilterRequestXLanguage) UnmarshalJSON(b []byte) error {
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
