package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGaussDbDatastoresRequest Request Object
type ListGaussDbDatastoresRequest struct {

	// 语言
	XLanguage *ListGaussDbDatastoresRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListGaussDbDatastoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussDbDatastoresRequest struct{}"
	}

	return strings.Join([]string{"ListGaussDbDatastoresRequest", string(data)}, " ")
}

type ListGaussDbDatastoresRequestXLanguage struct {
	value string
}

type ListGaussDbDatastoresRequestXLanguageEnum struct {
	ZH_CN ListGaussDbDatastoresRequestXLanguage
	EN_US ListGaussDbDatastoresRequestXLanguage
}

func GetListGaussDbDatastoresRequestXLanguageEnum() ListGaussDbDatastoresRequestXLanguageEnum {
	return ListGaussDbDatastoresRequestXLanguageEnum{
		ZH_CN: ListGaussDbDatastoresRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListGaussDbDatastoresRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListGaussDbDatastoresRequestXLanguage) Value() string {
	return c.value
}

func (c ListGaussDbDatastoresRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGaussDbDatastoresRequestXLanguage) UnmarshalJSON(b []byte) error {
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
