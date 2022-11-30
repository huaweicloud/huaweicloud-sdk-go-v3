package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListHistoryOperationsRequest struct {

	// 语言, 默认值为en-us。
	XLanguage *ListHistoryOperationsRequestXLanguage `json:"X-Language,omitempty"`

	// 参数模板ID。
	ConfigId string `json:"config_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHistoryOperationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryOperationsRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryOperationsRequest", string(data)}, " ")
}

type ListHistoryOperationsRequestXLanguage struct {
	value string
}

type ListHistoryOperationsRequestXLanguageEnum struct {
	ZH_CN ListHistoryOperationsRequestXLanguage
	EN_US ListHistoryOperationsRequestXLanguage
}

func GetListHistoryOperationsRequestXLanguageEnum() ListHistoryOperationsRequestXLanguageEnum {
	return ListHistoryOperationsRequestXLanguageEnum{
		ZH_CN: ListHistoryOperationsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListHistoryOperationsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListHistoryOperationsRequestXLanguage) Value() string {
	return c.value
}

func (c ListHistoryOperationsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHistoryOperationsRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
