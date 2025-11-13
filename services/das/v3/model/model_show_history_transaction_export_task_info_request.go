package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHistoryTransactionExportTaskInfoRequest Request Object
type ShowHistoryTransactionExportTaskInfoRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ShowHistoryTransactionExportTaskInfoRequestXLanguage `json:"X-Language,omitempty"`

	// 导出任务id
	TaskId int64 `json:"task_id"`
}

func (o ShowHistoryTransactionExportTaskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTransactionExportTaskInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowHistoryTransactionExportTaskInfoRequest", string(data)}, " ")
}

type ShowHistoryTransactionExportTaskInfoRequestXLanguage struct {
	value string
}

type ShowHistoryTransactionExportTaskInfoRequestXLanguageEnum struct {
	EN_US ShowHistoryTransactionExportTaskInfoRequestXLanguage
	ZH_CN ShowHistoryTransactionExportTaskInfoRequestXLanguage
}

func GetShowHistoryTransactionExportTaskInfoRequestXLanguageEnum() ShowHistoryTransactionExportTaskInfoRequestXLanguageEnum {
	return ShowHistoryTransactionExportTaskInfoRequestXLanguageEnum{
		EN_US: ShowHistoryTransactionExportTaskInfoRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowHistoryTransactionExportTaskInfoRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowHistoryTransactionExportTaskInfoRequestXLanguage) Value() string {
	return c.value
}

func (c ShowHistoryTransactionExportTaskInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHistoryTransactionExportTaskInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
