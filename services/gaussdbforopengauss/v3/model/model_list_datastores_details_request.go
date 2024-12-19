package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDatastoresDetailsRequest Request Object
type ListDatastoresDetailsRequest struct {

	// 语言
	XLanguage *ListDatastoresDetailsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListDatastoresDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoresDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListDatastoresDetailsRequest", string(data)}, " ")
}

type ListDatastoresDetailsRequestXLanguage struct {
	value string
}

type ListDatastoresDetailsRequestXLanguageEnum struct {
	ZH_CN ListDatastoresDetailsRequestXLanguage
	EN_US ListDatastoresDetailsRequestXLanguage
}

func GetListDatastoresDetailsRequestXLanguageEnum() ListDatastoresDetailsRequestXLanguageEnum {
	return ListDatastoresDetailsRequestXLanguageEnum{
		ZH_CN: ListDatastoresDetailsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListDatastoresDetailsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListDatastoresDetailsRequestXLanguage) Value() string {
	return c.value
}

func (c ListDatastoresDetailsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatastoresDetailsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
