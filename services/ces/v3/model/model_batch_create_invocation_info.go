package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchCreateInvocationInfo struct {

	// **参数解释**: 机器id **取值范围**: 1到64个字符的字符串，且只包含字母、数字和连字符
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 任务id **取值范围**: 1到64个字符的字符串，且只包含字母、数字和连字符
	InvocationId *string `json:"invocation_id,omitempty"`

	// **参数解释**: 任务结果, successful成功，error失败 **取值范围**: - successful: 成功 - error: 失败
	RetStatus *BatchCreateInvocationInfoRetStatus `json:"ret_status,omitempty"`

	// **参数解释**: 错误码 **取值范围**: 以\"invocationmgr.\"开头且后跟4位数字的完整字符串
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**: 错误信息 **取值范围**: 数组长度范围为[1,128]
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BatchCreateInvocationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInvocationInfo struct{}"
	}

	return strings.Join([]string{"BatchCreateInvocationInfo", string(data)}, " ")
}

type BatchCreateInvocationInfoRetStatus struct {
	value string
}

type BatchCreateInvocationInfoRetStatusEnum struct {
	SUCCESSFUL BatchCreateInvocationInfoRetStatus
	ERROR      BatchCreateInvocationInfoRetStatus
}

func GetBatchCreateInvocationInfoRetStatusEnum() BatchCreateInvocationInfoRetStatusEnum {
	return BatchCreateInvocationInfoRetStatusEnum{
		SUCCESSFUL: BatchCreateInvocationInfoRetStatus{
			value: "successful",
		},
		ERROR: BatchCreateInvocationInfoRetStatus{
			value: "error",
		},
	}
}

func (c BatchCreateInvocationInfoRetStatus) Value() string {
	return c.value
}

func (c BatchCreateInvocationInfoRetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateInvocationInfoRetStatus) UnmarshalJSON(b []byte) error {
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
