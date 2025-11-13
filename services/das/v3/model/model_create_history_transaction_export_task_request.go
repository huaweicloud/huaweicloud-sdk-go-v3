package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHistoryTransactionExportTaskRequest Request Object
type CreateHistoryTransactionExportTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *CreateHistoryTransactionExportTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateExportTaskReq `json:"body,omitempty"`
}

func (o CreateHistoryTransactionExportTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHistoryTransactionExportTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateHistoryTransactionExportTaskRequest", string(data)}, " ")
}

type CreateHistoryTransactionExportTaskRequestXLanguage struct {
	value string
}

type CreateHistoryTransactionExportTaskRequestXLanguageEnum struct {
	EN_US CreateHistoryTransactionExportTaskRequestXLanguage
	ZH_CN CreateHistoryTransactionExportTaskRequestXLanguage
}

func GetCreateHistoryTransactionExportTaskRequestXLanguageEnum() CreateHistoryTransactionExportTaskRequestXLanguageEnum {
	return CreateHistoryTransactionExportTaskRequestXLanguageEnum{
		EN_US: CreateHistoryTransactionExportTaskRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateHistoryTransactionExportTaskRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateHistoryTransactionExportTaskRequestXLanguage) Value() string {
	return c.value
}

func (c CreateHistoryTransactionExportTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHistoryTransactionExportTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
