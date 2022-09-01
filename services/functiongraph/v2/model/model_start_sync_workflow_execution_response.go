package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type StartSyncWorkflowExecutionResponse struct {

	// 流程实例ID
	ExecutionId *string `json:"execution_id,omitempty" xml:"execution_id"`

	// 流程执行最终状态
	Status *StartSyncWorkflowExecutionResponseStatus `json:"status,omitempty" xml:"status"`

	// 函数流的执行结果，JSON格式，仅在status为success时有值
	Output *interface{} `json:"output,omitempty" xml:"output"`

	// 流程执行错误信息，仅在status为fail时有值
	Errors *[]SyncExecutionNodeErrorDetail `json:"errors,omitempty" xml:"errors"`

	// 流程实例创建时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 流程实例结束时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	EndTime        *string `json:"end_time,omitempty" xml:"end_time"`
	HttpStatusCode int     `json:"-"`
}

func (o StartSyncWorkflowExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSyncWorkflowExecutionResponse struct{}"
	}

	return strings.Join([]string{"StartSyncWorkflowExecutionResponse", string(data)}, " ")
}

type StartSyncWorkflowExecutionResponseStatus struct {
	value string
}

type StartSyncWorkflowExecutionResponseStatusEnum struct {
	SUCCESS StartSyncWorkflowExecutionResponseStatus
	FAIL    StartSyncWorkflowExecutionResponseStatus
	TIMEOUT StartSyncWorkflowExecutionResponseStatus
}

func GetStartSyncWorkflowExecutionResponseStatusEnum() StartSyncWorkflowExecutionResponseStatusEnum {
	return StartSyncWorkflowExecutionResponseStatusEnum{
		SUCCESS: StartSyncWorkflowExecutionResponseStatus{
			value: "success",
		},
		FAIL: StartSyncWorkflowExecutionResponseStatus{
			value: "fail",
		},
		TIMEOUT: StartSyncWorkflowExecutionResponseStatus{
			value: "timeout",
		},
	}
}

func (c StartSyncWorkflowExecutionResponseStatus) Value() string {
	return c.value
}

func (c StartSyncWorkflowExecutionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartSyncWorkflowExecutionResponseStatus) UnmarshalJSON(b []byte) error {
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
