package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RemediationExecutionStatus 指定资源合规规则修正最新执行结果。
type RemediationExecutionStatus struct {
	ResourceKey *RemediationResourceKey `json:"resource_key,omitempty"`

	// 补救执行的开始时间。
	InvocationTime *string `json:"invocation_time,omitempty"`

	// 合规规则修正执行的状态。
	State *RemediationExecutionStatusState `json:"state,omitempty"`

	// 合规规则修正执行的信息。
	Message *string `json:"message,omitempty"`
}

func (o RemediationExecutionStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemediationExecutionStatus struct{}"
	}

	return strings.Join([]string{"RemediationExecutionStatus", string(data)}, " ")
}

type RemediationExecutionStatusState struct {
	value string
}

type RemediationExecutionStatusStateEnum struct {
	IN_QUEUE    RemediationExecutionStatusState
	IN_PROGRESS RemediationExecutionStatusState
	SUCCEEDED   RemediationExecutionStatusState
	FAILED      RemediationExecutionStatusState
}

func GetRemediationExecutionStatusStateEnum() RemediationExecutionStatusStateEnum {
	return RemediationExecutionStatusStateEnum{
		IN_QUEUE: RemediationExecutionStatusState{
			value: "IN_QUEUE",
		},
		IN_PROGRESS: RemediationExecutionStatusState{
			value: "IN_PROGRESS",
		},
		SUCCEEDED: RemediationExecutionStatusState{
			value: "SUCCEEDED",
		},
		FAILED: RemediationExecutionStatusState{
			value: "FAILED",
		},
	}
}

func (c RemediationExecutionStatusState) Value() string {
	return c.value
}

func (c RemediationExecutionStatusState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RemediationExecutionStatusState) UnmarshalJSON(b []byte) error {
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
