package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListParameterGroupTemplatesRequest Request Object
type ListParameterGroupTemplatesRequest struct {

	// 语言
	XLanguage *ListParameterGroupTemplatesRequestXLanguage `json:"X-Language,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListParameterGroupTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListParameterGroupTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListParameterGroupTemplatesRequest", string(data)}, " ")
}

type ListParameterGroupTemplatesRequestXLanguage struct {
	value string
}

type ListParameterGroupTemplatesRequestXLanguageEnum struct {
	ZH_CN ListParameterGroupTemplatesRequestXLanguage
	EN_US ListParameterGroupTemplatesRequestXLanguage
}

func GetListParameterGroupTemplatesRequestXLanguageEnum() ListParameterGroupTemplatesRequestXLanguageEnum {
	return ListParameterGroupTemplatesRequestXLanguageEnum{
		ZH_CN: ListParameterGroupTemplatesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListParameterGroupTemplatesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListParameterGroupTemplatesRequestXLanguage) Value() string {
	return c.value
}

func (c ListParameterGroupTemplatesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListParameterGroupTemplatesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
