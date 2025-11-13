package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHistoryTransactionExportTaskRequest Request Object
type ListHistoryTransactionExportTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ListHistoryTransactionExportTaskRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询,默认为0(偏移0条数据,表示从第一条数据开始查询),必须为数字,不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100,不能为负数,最小值为1,最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHistoryTransactionExportTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryTransactionExportTaskRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryTransactionExportTaskRequest", string(data)}, " ")
}

type ListHistoryTransactionExportTaskRequestXLanguage struct {
	value string
}

type ListHistoryTransactionExportTaskRequestXLanguageEnum struct {
	EN_US ListHistoryTransactionExportTaskRequestXLanguage
	ZH_CN ListHistoryTransactionExportTaskRequestXLanguage
}

func GetListHistoryTransactionExportTaskRequestXLanguageEnum() ListHistoryTransactionExportTaskRequestXLanguageEnum {
	return ListHistoryTransactionExportTaskRequestXLanguageEnum{
		EN_US: ListHistoryTransactionExportTaskRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListHistoryTransactionExportTaskRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListHistoryTransactionExportTaskRequestXLanguage) Value() string {
	return c.value
}

func (c ListHistoryTransactionExportTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHistoryTransactionExportTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
