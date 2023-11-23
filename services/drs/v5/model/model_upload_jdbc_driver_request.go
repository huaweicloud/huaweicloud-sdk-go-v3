package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UploadJdbcDriverRequest Request Object
type UploadJdbcDriverRequest struct {

	// 请求语言类型。
	XLanguage *UploadJdbcDriverRequestXLanguage `json:"X-Language,omitempty"`

	Body *UploadJdbcDriverRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadJdbcDriverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadJdbcDriverRequest struct{}"
	}

	return strings.Join([]string{"UploadJdbcDriverRequest", string(data)}, " ")
}

type UploadJdbcDriverRequestXLanguage struct {
	value string
}

type UploadJdbcDriverRequestXLanguageEnum struct {
	EN_US UploadJdbcDriverRequestXLanguage
	ZH_CN UploadJdbcDriverRequestXLanguage
}

func GetUploadJdbcDriverRequestXLanguageEnum() UploadJdbcDriverRequestXLanguageEnum {
	return UploadJdbcDriverRequestXLanguageEnum{
		EN_US: UploadJdbcDriverRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UploadJdbcDriverRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UploadJdbcDriverRequestXLanguage) Value() string {
	return c.value
}

func (c UploadJdbcDriverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadJdbcDriverRequestXLanguage) UnmarshalJSON(b []byte) error {
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
