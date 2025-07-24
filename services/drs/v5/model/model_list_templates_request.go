package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTemplatesRequest Request Object
type ListTemplatesRequest struct {

	// 请求语言类型。
	XLanguage *ListTemplatesRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}

type ListTemplatesRequestXLanguage struct {
	value string
}

type ListTemplatesRequestXLanguageEnum struct {
	EN_US ListTemplatesRequestXLanguage
	ZH_CN ListTemplatesRequestXLanguage
}

func GetListTemplatesRequestXLanguageEnum() ListTemplatesRequestXLanguageEnum {
	return ListTemplatesRequestXLanguageEnum{
		EN_US: ListTemplatesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListTemplatesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListTemplatesRequestXLanguage) Value() string {
	return c.value
}

func (c ListTemplatesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTemplatesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
