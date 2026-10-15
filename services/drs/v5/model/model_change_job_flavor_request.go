package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeJobFlavorRequest Request Object
type ChangeJobFlavorRequest struct {

	// 规格变更任务ID。示例值：27f6a81-7d6b-4d56-a184-7d30576jb201
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ChangeJobFlavorRequestXLanguage `json:"X-Language,omitempty"`

	Body *ChangeFlavorReq `json:"body,omitempty"`
}

func (o ChangeJobFlavorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeJobFlavorRequest struct{}"
	}

	return strings.Join([]string{"ChangeJobFlavorRequest", string(data)}, " ")
}

type ChangeJobFlavorRequestXLanguage struct {
	value string
}

type ChangeJobFlavorRequestXLanguageEnum struct {
	EN_US ChangeJobFlavorRequestXLanguage
	ZH_CN ChangeJobFlavorRequestXLanguage
	ES_US ChangeJobFlavorRequestXLanguage
	PT_BR ChangeJobFlavorRequestXLanguage
}

func GetChangeJobFlavorRequestXLanguageEnum() ChangeJobFlavorRequestXLanguageEnum {
	return ChangeJobFlavorRequestXLanguageEnum{
		EN_US: ChangeJobFlavorRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ChangeJobFlavorRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: ChangeJobFlavorRequestXLanguage{
			value: "es-us",
		},
		PT_BR: ChangeJobFlavorRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c ChangeJobFlavorRequestXLanguage) Value() string {
	return c.value
}

func (c ChangeJobFlavorRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeJobFlavorRequestXLanguage) UnmarshalJSON(b []byte) error {
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
