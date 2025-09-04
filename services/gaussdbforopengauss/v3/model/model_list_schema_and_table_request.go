package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSchemaAndTableRequest Request Object
type ListSchemaAndTableRequest struct {

	// 语言, 默认值为en-us。
	XLanguage *ListSchemaAndTableRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ListSchemaAndTableRequestBody `json:"body,omitempty"`
}

func (o ListSchemaAndTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemaAndTableRequest struct{}"
	}

	return strings.Join([]string{"ListSchemaAndTableRequest", string(data)}, " ")
}

type ListSchemaAndTableRequestXLanguage struct {
	value string
}

type ListSchemaAndTableRequestXLanguageEnum struct {
	ZH_CN ListSchemaAndTableRequestXLanguage
	EN_US ListSchemaAndTableRequestXLanguage
}

func GetListSchemaAndTableRequestXLanguageEnum() ListSchemaAndTableRequestXLanguageEnum {
	return ListSchemaAndTableRequestXLanguageEnum{
		ZH_CN: ListSchemaAndTableRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSchemaAndTableRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSchemaAndTableRequestXLanguage) Value() string {
	return c.value
}

func (c ListSchemaAndTableRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSchemaAndTableRequestXLanguage) UnmarshalJSON(b []byte) error {
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
