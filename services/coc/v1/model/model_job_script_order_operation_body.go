package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobScriptOrderOperationBody 工单操作对象 operation_type为 CANCEL_INSTANCE：取消实例，instance_id 不能为空； SKIP_BATCH： 跳过当前批次，batch_index不能为空； CANCEL_ORDER：取消工单，batch_index和instance_id可为空； PAUSE_ORDER：暂停工单，batch_index和instance_id可为空； CONTINUE_ORDER：继续工单，batch_index和instance_id可为空
type JobScriptOrderOperationBody struct {

	// 适用于批次操作时
	BatchIndex *int32 `json:"batch_index,omitempty"`

	// 适用于实例操作时 非script_execute_instance_do中iD，需要新增字段
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 操作类型：取消实例、跳过批次、取消整个工单、暂停整个工单、继续整个工单 CANCEL_INSTANCE：取消实例 SKIP_BATCH：跳过批次 CANCEL_ORDER：取消整个工单 PAUSE_ORDER：暂停整个工单 CONTINUE_ORDER：继续整个工单
	OperationType JobScriptOrderOperationBodyOperationType `json:"operation_type"`
}

func (o JobScriptOrderOperationBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScriptOrderOperationBody struct{}"
	}

	return strings.Join([]string{"JobScriptOrderOperationBody", string(data)}, " ")
}

type JobScriptOrderOperationBodyOperationType struct {
	value string
}

type JobScriptOrderOperationBodyOperationTypeEnum struct {
	CANCEL_INSTANCE JobScriptOrderOperationBodyOperationType
	SKIP_BATCH      JobScriptOrderOperationBodyOperationType
	CANCEL_ORDER    JobScriptOrderOperationBodyOperationType
	PAUSE_ORDER     JobScriptOrderOperationBodyOperationType
	CONTINUE_ORDER  JobScriptOrderOperationBodyOperationType
}

func GetJobScriptOrderOperationBodyOperationTypeEnum() JobScriptOrderOperationBodyOperationTypeEnum {
	return JobScriptOrderOperationBodyOperationTypeEnum{
		CANCEL_INSTANCE: JobScriptOrderOperationBodyOperationType{
			value: "CANCEL_INSTANCE",
		},
		SKIP_BATCH: JobScriptOrderOperationBodyOperationType{
			value: "SKIP_BATCH",
		},
		CANCEL_ORDER: JobScriptOrderOperationBodyOperationType{
			value: "CANCEL_ORDER",
		},
		PAUSE_ORDER: JobScriptOrderOperationBodyOperationType{
			value: "PAUSE_ORDER",
		},
		CONTINUE_ORDER: JobScriptOrderOperationBodyOperationType{
			value: "CONTINUE_ORDER",
		},
	}
}

func (c JobScriptOrderOperationBodyOperationType) Value() string {
	return c.value
}

func (c JobScriptOrderOperationBodyOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobScriptOrderOperationBodyOperationType) UnmarshalJSON(b []byte) error {
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
