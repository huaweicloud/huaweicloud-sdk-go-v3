package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteHistoryTransactionExportTaskRequest Request Object
type DeleteHistoryTransactionExportTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *DeleteHistoryTransactionExportTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateExportTaskResp `json:"body,omitempty"`
}

func (o DeleteHistoryTransactionExportTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHistoryTransactionExportTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteHistoryTransactionExportTaskRequest", string(data)}, " ")
}

type DeleteHistoryTransactionExportTaskRequestXLanguage struct {
	value string
}

type DeleteHistoryTransactionExportTaskRequestXLanguageEnum struct {
	EN_US DeleteHistoryTransactionExportTaskRequestXLanguage
	ZH_CN DeleteHistoryTransactionExportTaskRequestXLanguage
}

func GetDeleteHistoryTransactionExportTaskRequestXLanguageEnum() DeleteHistoryTransactionExportTaskRequestXLanguageEnum {
	return DeleteHistoryTransactionExportTaskRequestXLanguageEnum{
		EN_US: DeleteHistoryTransactionExportTaskRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteHistoryTransactionExportTaskRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteHistoryTransactionExportTaskRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteHistoryTransactionExportTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteHistoryTransactionExportTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
