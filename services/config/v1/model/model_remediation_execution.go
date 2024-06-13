package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RemediationExecution 合规规则修正执行结果的详情。
type RemediationExecution struct {

	// 是否为自动修正。
	Automatic *bool `json:"automatic,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 云服务名称。
	ResourceProvider *string `json:"resource_provider,omitempty"`

	// 资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 补救执行的开始时间。
	InvocationTime *string `json:"invocation_time,omitempty"`

	// 合规规则修正执行的状态。
	State *RemediationExecutionState `json:"state,omitempty"`

	// 合规规则修正执行的信息。
	Message *string `json:"message,omitempty"`
}

func (o RemediationExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemediationExecution struct{}"
	}

	return strings.Join([]string{"RemediationExecution", string(data)}, " ")
}

type RemediationExecutionState struct {
	value string
}

type RemediationExecutionStateEnum struct {
	IN_QUEUE    RemediationExecutionState
	IN_PROGRESS RemediationExecutionState
	SUCCEEDED   RemediationExecutionState
	FAILED      RemediationExecutionState
}

func GetRemediationExecutionStateEnum() RemediationExecutionStateEnum {
	return RemediationExecutionStateEnum{
		IN_QUEUE: RemediationExecutionState{
			value: "IN_QUEUE",
		},
		IN_PROGRESS: RemediationExecutionState{
			value: "IN_PROGRESS",
		},
		SUCCEEDED: RemediationExecutionState{
			value: "SUCCEEDED",
		},
		FAILED: RemediationExecutionState{
			value: "FAILED",
		},
	}
}

func (c RemediationExecutionState) Value() string {
	return c.value
}

func (c RemediationExecutionState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RemediationExecutionState) UnmarshalJSON(b []byte) error {
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
