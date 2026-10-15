package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListReplaySupportSourceConfigRequest Request Object
type ListReplaySupportSourceConfigRequest struct {

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`

	// 请求语言类型。
	XLanguage *ListReplaySupportSourceConfigRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListReplaySupportSourceConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplaySupportSourceConfigRequest struct{}"
	}

	return strings.Join([]string{"ListReplaySupportSourceConfigRequest", string(data)}, " ")
}

type ListReplaySupportSourceConfigRequestXLanguage struct {
	value string
}

type ListReplaySupportSourceConfigRequestXLanguageEnum struct {
	EN_US ListReplaySupportSourceConfigRequestXLanguage
	ZH_CN ListReplaySupportSourceConfigRequestXLanguage
}

func GetListReplaySupportSourceConfigRequestXLanguageEnum() ListReplaySupportSourceConfigRequestXLanguageEnum {
	return ListReplaySupportSourceConfigRequestXLanguageEnum{
		EN_US: ListReplaySupportSourceConfigRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListReplaySupportSourceConfigRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListReplaySupportSourceConfigRequestXLanguage) Value() string {
	return c.value
}

func (c ListReplaySupportSourceConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListReplaySupportSourceConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
