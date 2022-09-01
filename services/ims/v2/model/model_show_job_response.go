package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowJobResponse struct {

	// 任务状态，目前取值如下： SUCCESS：表示该任务执行已经结束，任务执行成功。 FAIL：表示该任务执行已经结束，任务执行失败。 RUNNING：表示该任务正在执行。 INIT：表给任务还未执行，正在初始化。
	Status *ShowJobResponseStatus `json:"status,omitempty" xml:"status"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 任务类型。
	JobType *string `json:"job_type,omitempty" xml:"job_type"`

	// 任务开始执行时间。格式为UTC时间。
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 任务结束时间。格式为UTC时间。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 失败原因。
	FailReason *string `json:"fail_reason,omitempty" xml:"fail_reason"`

	Entities       *JobEntities `json:"entities,omitempty" xml:"entities"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}

type ShowJobResponseStatus struct {
	value string
}

type ShowJobResponseStatusEnum struct {
	SUCCESS ShowJobResponseStatus
	FAIL    ShowJobResponseStatus
	RUNNING ShowJobResponseStatus
	INIT    ShowJobResponseStatus
}

func GetShowJobResponseStatusEnum() ShowJobResponseStatusEnum {
	return ShowJobResponseStatusEnum{
		SUCCESS: ShowJobResponseStatus{
			value: "SUCCESS",
		},
		FAIL: ShowJobResponseStatus{
			value: "FAIL",
		},
		RUNNING: ShowJobResponseStatus{
			value: "RUNNING",
		},
		INIT: ShowJobResponseStatus{
			value: "INIT",
		},
	}
}

func (c ShowJobResponseStatus) Value() string {
	return c.value
}

func (c ShowJobResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobResponseStatus) UnmarshalJSON(b []byte) error {
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
