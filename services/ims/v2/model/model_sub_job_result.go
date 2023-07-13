package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SubJobResult struct {

	// 任务状态，目前取值如下： SUCCESS：表示该任务执行已经结束，任务执行成功。 FAIL：表示该任务执行已经结束，任务执行失败。 RUNNING：表示该任务正在执行。 INIT：表给任务还未执行，正在初始化。
	Status *SubJobResultStatus `json:"status,omitempty"`

	// 子任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 子任务类型。
	JobType *string `json:"job_type,omitempty"`

	// 子任务开始执行时间。格式为UTC时间。
	BeginTime *string `json:"begin_time,omitempty"`

	// 子任务结束时间。格式为UTC时间。
	EndTime *string `json:"end_time,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 失败原因。
	FailReason *string `json:"fail_reason,omitempty"`

	Entities *SubJobEntities `json:"entities,omitempty"`
}

func (o SubJobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobResult struct{}"
	}

	return strings.Join([]string{"SubJobResult", string(data)}, " ")
}

type SubJobResultStatus struct {
	value string
}

type SubJobResultStatusEnum struct {
	SUCCESS SubJobResultStatus
	FAIL    SubJobResultStatus
	RUNNING SubJobResultStatus
	INIT    SubJobResultStatus
}

func GetSubJobResultStatusEnum() SubJobResultStatusEnum {
	return SubJobResultStatusEnum{
		SUCCESS: SubJobResultStatus{
			value: "SUCCESS",
		},
		FAIL: SubJobResultStatus{
			value: "FAIL",
		},
		RUNNING: SubJobResultStatus{
			value: "RUNNING",
		},
		INIT: SubJobResultStatus{
			value: "INIT",
		},
	}
}

func (c SubJobResultStatus) Value() string {
	return c.value
}

func (c SubJobResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubJobResultStatus) UnmarshalJSON(b []byte) error {
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
