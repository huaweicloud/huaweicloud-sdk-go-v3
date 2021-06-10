package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RefreshPreheatingBody struct {
	// 任务id。

	Id *string `json:"id,omitempty"`
	// 任务的类型， 其值可以为refresh或preheating。

	TaskType *RefreshPreheatingBodyTaskType `json:"task_type,omitempty"`
	// 刷新结果。task_done表示刷新成功  ，task_inprocess表示刷新中。

	Status *RefreshPreheatingBodyStatus `json:"status,omitempty"`
	// 处理中的url个数。

	Processing *int32 `json:"processing,omitempty"`
	// 成功处理的url个数。

	Succeed *int32 `json:"succeed,omitempty"`
	// 处理失败的url个数。

	Failed *int32 `json:"failed,omitempty"`
	// 总共的任务个数。

	Total *int32 `json:"total,omitempty"`
	// 任务创建时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。

	CreateTime *int64 `json:"create_time,omitempty"`
	// 刷新缓存的url列表。

	Urls *[]string `json:"urls,omitempty"`
}

func (o RefreshPreheatingBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RefreshPreheatingBody struct{}"
	}

	return strings.Join([]string{"RefreshPreheatingBody", string(data)}, " ")
}

type RefreshPreheatingBodyTaskType struct {
	value string
}

type RefreshPreheatingBodyTaskTypeEnum struct {
	REFRESH    RefreshPreheatingBodyTaskType
	PREHEATING RefreshPreheatingBodyTaskType
}

func GetRefreshPreheatingBodyTaskTypeEnum() RefreshPreheatingBodyTaskTypeEnum {
	return RefreshPreheatingBodyTaskTypeEnum{
		REFRESH: RefreshPreheatingBodyTaskType{
			value: "refresh",
		},
		PREHEATING: RefreshPreheatingBodyTaskType{
			value: "preheating",
		},
	}
}

func (c RefreshPreheatingBodyTaskType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *RefreshPreheatingBodyTaskType) UnmarshalJSON(b []byte) error {
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

type RefreshPreheatingBodyStatus struct {
	value string
}

type RefreshPreheatingBodyStatusEnum struct {
	TASK_DONE RefreshPreheatingBodyStatus
}

func GetRefreshPreheatingBodyStatusEnum() RefreshPreheatingBodyStatusEnum {
	return RefreshPreheatingBodyStatusEnum{
		TASK_DONE: RefreshPreheatingBodyStatus{
			value: "task_done",
		},
	}
}

func (c RefreshPreheatingBodyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *RefreshPreheatingBodyStatus) UnmarshalJSON(b []byte) error {
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
