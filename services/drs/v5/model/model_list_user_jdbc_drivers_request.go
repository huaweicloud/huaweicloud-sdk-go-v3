package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListUserJdbcDriversRequest Request Object
type ListUserJdbcDriversRequest struct {

	// 每页显示的条目数量。默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset 大于等于 0。默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定待查询的驱动文件类型。取值范围： - db2：DB2 for LUW - informix：Informix
	DriverType ListUserJdbcDriversRequestDriverType `json:"driver_type"`

	// 请求语言类型。
	XLanguage *ListUserJdbcDriversRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListUserJdbcDriversRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserJdbcDriversRequest struct{}"
	}

	return strings.Join([]string{"ListUserJdbcDriversRequest", string(data)}, " ")
}

type ListUserJdbcDriversRequestDriverType struct {
	value string
}

type ListUserJdbcDriversRequestDriverTypeEnum struct {
	DB2      ListUserJdbcDriversRequestDriverType
	INFORMIX ListUserJdbcDriversRequestDriverType
}

func GetListUserJdbcDriversRequestDriverTypeEnum() ListUserJdbcDriversRequestDriverTypeEnum {
	return ListUserJdbcDriversRequestDriverTypeEnum{
		DB2: ListUserJdbcDriversRequestDriverType{
			value: "db2",
		},
		INFORMIX: ListUserJdbcDriversRequestDriverType{
			value: "informix",
		},
	}
}

func (c ListUserJdbcDriversRequestDriverType) Value() string {
	return c.value
}

func (c ListUserJdbcDriversRequestDriverType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUserJdbcDriversRequestDriverType) UnmarshalJSON(b []byte) error {
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

type ListUserJdbcDriversRequestXLanguage struct {
	value string
}

type ListUserJdbcDriversRequestXLanguageEnum struct {
	EN_US ListUserJdbcDriversRequestXLanguage
	ZH_CN ListUserJdbcDriversRequestXLanguage
}

func GetListUserJdbcDriversRequestXLanguageEnum() ListUserJdbcDriversRequestXLanguageEnum {
	return ListUserJdbcDriversRequestXLanguageEnum{
		EN_US: ListUserJdbcDriversRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListUserJdbcDriversRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListUserJdbcDriversRequestXLanguage) Value() string {
	return c.value
}

func (c ListUserJdbcDriversRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUserJdbcDriversRequestXLanguage) UnmarshalJSON(b []byte) error {
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
