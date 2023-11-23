package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListJdbcDriversRequest Request Object
type ListJdbcDriversRequest struct {

	// 每页显示的条目数量。默认为10
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset 大于等于 0。默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 请求语言类型。
	XLanguage *ListJdbcDriversRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListJdbcDriversRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJdbcDriversRequest struct{}"
	}

	return strings.Join([]string{"ListJdbcDriversRequest", string(data)}, " ")
}

type ListJdbcDriversRequestXLanguage struct {
	value string
}

type ListJdbcDriversRequestXLanguageEnum struct {
	EN_US ListJdbcDriversRequestXLanguage
	ZH_CN ListJdbcDriversRequestXLanguage
}

func GetListJdbcDriversRequestXLanguageEnum() ListJdbcDriversRequestXLanguageEnum {
	return ListJdbcDriversRequestXLanguageEnum{
		EN_US: ListJdbcDriversRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListJdbcDriversRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListJdbcDriversRequestXLanguage) Value() string {
	return c.value
}

func (c ListJdbcDriversRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJdbcDriversRequestXLanguage) UnmarshalJSON(b []byte) error {
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
