package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 异步调用记录
type ListFunctionAsyncInvocationsResult struct {

	// 异步调用请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 异步调用状态，支持5种状态 WAIT: 等待 RUNNING: 执行中 SUCCESS: 执行成功 FAIL: 执行失败 DISCARD: 请求丢弃
	Status *ListFunctionAsyncInvocationsResultStatus `json:"status,omitempty"`

	// 异步调用错误信息，如果执行成功，则返回空
	ErrorMessage *string `json:"error_message,omitempty"`

	// 异步调用开始时间（格式为YYYY-MM-DD'T'HH:mm:ss,UTC时间）。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 异步调用结束时间（格式为YYYY-MM-DD'T'HH:mm:ss,UTC时间）。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`
}

func (o ListFunctionAsyncInvocationsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionAsyncInvocationsResult struct{}"
	}

	return strings.Join([]string{"ListFunctionAsyncInvocationsResult", string(data)}, " ")
}

type ListFunctionAsyncInvocationsResultStatus struct {
	value string
}

type ListFunctionAsyncInvocationsResultStatusEnum struct {
	WAIT    ListFunctionAsyncInvocationsResultStatus
	RUNNING ListFunctionAsyncInvocationsResultStatus
	SUCCESS ListFunctionAsyncInvocationsResultStatus
	FAIL    ListFunctionAsyncInvocationsResultStatus
	DISCARD ListFunctionAsyncInvocationsResultStatus
}

func GetListFunctionAsyncInvocationsResultStatusEnum() ListFunctionAsyncInvocationsResultStatusEnum {
	return ListFunctionAsyncInvocationsResultStatusEnum{
		WAIT: ListFunctionAsyncInvocationsResultStatus{
			value: "WAIT",
		},
		RUNNING: ListFunctionAsyncInvocationsResultStatus{
			value: "RUNNING",
		},
		SUCCESS: ListFunctionAsyncInvocationsResultStatus{
			value: "SUCCESS",
		},
		FAIL: ListFunctionAsyncInvocationsResultStatus{
			value: "FAIL",
		},
		DISCARD: ListFunctionAsyncInvocationsResultStatus{
			value: "DISCARD",
		},
	}
}

func (c ListFunctionAsyncInvocationsResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFunctionAsyncInvocationsResultStatus) UnmarshalJSON(b []byte) error {
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
