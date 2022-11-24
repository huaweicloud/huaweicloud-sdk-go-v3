package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type FlowExecutionBriefV2 struct {

	// 函数流ID
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 函数流执行ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 函数流执行状态
	Status *FlowExecutionBriefV2Status `json:"status,omitempty"`

	// 开始时间（格式为yyyy-MM-dd'T'HH:mm:ss'Z',UTC时间）。
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 结束时间（格式为yyyy-MM-dd'T'HH:mm:ss'Z',UTC时间）。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 结束时间（格式为yyyy-MM-dd'T'HH:mm:ss'Z',UTC时间）。
	LastUpdateTime *sdktime.SdkTime `json:"last_update_time,omitempty"`

	CreateBy *string `json:"create_by,omitempty"`

	// 函数流执行urn
	WorkflowUrn *string `json:"workflow_urn,omitempty"`
}

func (o FlowExecutionBriefV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowExecutionBriefV2 struct{}"
	}

	return strings.Join([]string{"FlowExecutionBriefV2", string(data)}, " ")
}

type FlowExecutionBriefV2Status struct {
	value string
}

type FlowExecutionBriefV2StatusEnum struct {
	SUCCESS   FlowExecutionBriefV2Status
	FAIL      FlowExecutionBriefV2Status
	RUNNING   FlowExecutionBriefV2Status
	TIMEOUT   FlowExecutionBriefV2Status
	CANCEL    FlowExecutionBriefV2Status
	SCHEDULED FlowExecutionBriefV2Status
	RECOVERED FlowExecutionBriefV2Status
	RETRYING  FlowExecutionBriefV2Status
	STOPPING  FlowExecutionBriefV2Status
}

func GetFlowExecutionBriefV2StatusEnum() FlowExecutionBriefV2StatusEnum {
	return FlowExecutionBriefV2StatusEnum{
		SUCCESS: FlowExecutionBriefV2Status{
			value: "success",
		},
		FAIL: FlowExecutionBriefV2Status{
			value: "fail",
		},
		RUNNING: FlowExecutionBriefV2Status{
			value: "running",
		},
		TIMEOUT: FlowExecutionBriefV2Status{
			value: "timeout",
		},
		CANCEL: FlowExecutionBriefV2Status{
			value: "cancel",
		},
		SCHEDULED: FlowExecutionBriefV2Status{
			value: "scheduled",
		},
		RECOVERED: FlowExecutionBriefV2Status{
			value: "recovered",
		},
		RETRYING: FlowExecutionBriefV2Status{
			value: "retrying",
		},
		STOPPING: FlowExecutionBriefV2Status{
			value: "stopping",
		},
	}
}

func (c FlowExecutionBriefV2Status) Value() string {
	return c.value
}

func (c FlowExecutionBriefV2Status) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowExecutionBriefV2Status) UnmarshalJSON(b []byte) error {
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
