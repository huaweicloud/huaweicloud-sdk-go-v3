package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CopyJobRequest Request Object
type CopyJobRequest struct {

	// 请求语言类型。
	XLanguage *CopyJobRequestXLanguage `json:"X-Language,omitempty"`

	Body *CloneJobReq `json:"body,omitempty"`
}

func (o CopyJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyJobRequest struct{}"
	}

	return strings.Join([]string{"CopyJobRequest", string(data)}, " ")
}

type CopyJobRequestXLanguage struct {
	value string
}

type CopyJobRequestXLanguageEnum struct {
	EN_US CopyJobRequestXLanguage
	ZH_CN CopyJobRequestXLanguage
}

func GetCopyJobRequestXLanguageEnum() CopyJobRequestXLanguageEnum {
	return CopyJobRequestXLanguageEnum{
		EN_US: CopyJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CopyJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CopyJobRequestXLanguage) Value() string {
	return c.value
}

func (c CopyJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CopyJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
