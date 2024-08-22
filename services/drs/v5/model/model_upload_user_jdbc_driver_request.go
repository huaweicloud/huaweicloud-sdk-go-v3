package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UploadUserJdbcDriverRequest Request Object
type UploadUserJdbcDriverRequest struct {

	// 请求语言类型。
	XLanguage *UploadUserJdbcDriverRequestXLanguage `json:"X-Language,omitempty"`

	// 指定待上传的驱动文件类型。取值范围： - db2：DB2 for LUW - informix：Informix
	DriverType UploadUserJdbcDriverRequestDriverType `json:"driver_type"`

	Body *UploadUserJdbcDriverRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadUserJdbcDriverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadUserJdbcDriverRequest struct{}"
	}

	return strings.Join([]string{"UploadUserJdbcDriverRequest", string(data)}, " ")
}

type UploadUserJdbcDriverRequestXLanguage struct {
	value string
}

type UploadUserJdbcDriverRequestXLanguageEnum struct {
	EN_US UploadUserJdbcDriverRequestXLanguage
	ZH_CN UploadUserJdbcDriverRequestXLanguage
}

func GetUploadUserJdbcDriverRequestXLanguageEnum() UploadUserJdbcDriverRequestXLanguageEnum {
	return UploadUserJdbcDriverRequestXLanguageEnum{
		EN_US: UploadUserJdbcDriverRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UploadUserJdbcDriverRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UploadUserJdbcDriverRequestXLanguage) Value() string {
	return c.value
}

func (c UploadUserJdbcDriverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadUserJdbcDriverRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type UploadUserJdbcDriverRequestDriverType struct {
	value string
}

type UploadUserJdbcDriverRequestDriverTypeEnum struct {
	DB2      UploadUserJdbcDriverRequestDriverType
	INFORMIX UploadUserJdbcDriverRequestDriverType
}

func GetUploadUserJdbcDriverRequestDriverTypeEnum() UploadUserJdbcDriverRequestDriverTypeEnum {
	return UploadUserJdbcDriverRequestDriverTypeEnum{
		DB2: UploadUserJdbcDriverRequestDriverType{
			value: "db2",
		},
		INFORMIX: UploadUserJdbcDriverRequestDriverType{
			value: "informix",
		},
	}
}

func (c UploadUserJdbcDriverRequestDriverType) Value() string {
	return c.value
}

func (c UploadUserJdbcDriverRequestDriverType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadUserJdbcDriverRequestDriverType) UnmarshalJSON(b []byte) error {
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
