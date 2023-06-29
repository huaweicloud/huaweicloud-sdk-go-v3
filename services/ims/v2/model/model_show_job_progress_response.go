package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowJobProgressResponse Response Object
type ShowJobProgressResponse struct {

	// 任务状态，目前取值如下： SUCCESS：表示该任务执行已经结束，任务执行成功。 FAIL：表示该任务执行已经结束，任务执行失败。 RUNNING：表示该任务正在执行。 INIT：表给任务还未执行，正在初始化。
	Status *ShowJobProgressResponseStatus `json:"status,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务类型。
	JobType *string `json:"job_type,omitempty"`

	// 任务开始执行时间。格式为UTC时间。
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间。格式为UTC时间。
	EndTime *string `json:"end_time,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 失败原因。
	FailReason *string `json:"fail_reason,omitempty"`

	Entities       *JobProgressEntities `json:"entities,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowJobProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowJobProgressResponse", string(data)}, " ")
}

type ShowJobProgressResponseStatus struct {
	value string
}

type ShowJobProgressResponseStatusEnum struct {
	SUCCESS ShowJobProgressResponseStatus
	FAIL    ShowJobProgressResponseStatus
	RUNNING ShowJobProgressResponseStatus
	INIT    ShowJobProgressResponseStatus
}

func GetShowJobProgressResponseStatusEnum() ShowJobProgressResponseStatusEnum {
	return ShowJobProgressResponseStatusEnum{
		SUCCESS: ShowJobProgressResponseStatus{
			value: "SUCCESS",
		},
		FAIL: ShowJobProgressResponseStatus{
			value: "FAIL",
		},
		RUNNING: ShowJobProgressResponseStatus{
			value: "RUNNING",
		},
		INIT: ShowJobProgressResponseStatus{
			value: "INIT",
		},
	}
}

func (c ShowJobProgressResponseStatus) Value() string {
	return c.value
}

func (c ShowJobProgressResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobProgressResponseStatus) UnmarshalJSON(b []byte) error {
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
