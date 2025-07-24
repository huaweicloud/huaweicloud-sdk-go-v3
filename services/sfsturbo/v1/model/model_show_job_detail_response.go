package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowJobDetailResponse Response Object
type ShowJobDetailResponse struct {

	// job的状态。success：成功。running：运行中。failed：失败。waiting：等待执行。
	Status *ShowJobDetailResponseStatus `json:"status,omitempty"`

	// job的ID。
	JobId *string `json:"job_id,omitempty"`

	// job的类型。
	JobType *string `json:"job_type,omitempty"`

	// job开始时间。UTC时间，格式：'2016-01-02 15:04:05'。
	BeginTime *string `json:"begin_time,omitempty"`

	// job结束时间。UTC时间，格式：'2016-01-02 15:04:05'。
	EndTime *string `json:"end_time,omitempty"`

	// job执行失败时的错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// job执行失败时的错误原因。
	FailReason *string `json:"fail_reason,omitempty"`

	// 子任务列表。
	SubJobs *[]GetSubJobDetail `json:"sub_jobs,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowJobDetailResponse", string(data)}, " ")
}

type ShowJobDetailResponseStatus struct {
	value string
}

type ShowJobDetailResponseStatusEnum struct {
	SUCCESS ShowJobDetailResponseStatus
	FAILED  ShowJobDetailResponseStatus
	WAITING ShowJobDetailResponseStatus
	RUNNING ShowJobDetailResponseStatus
}

func GetShowJobDetailResponseStatusEnum() ShowJobDetailResponseStatusEnum {
	return ShowJobDetailResponseStatusEnum{
		SUCCESS: ShowJobDetailResponseStatus{
			value: "success",
		},
		FAILED: ShowJobDetailResponseStatus{
			value: "failed",
		},
		WAITING: ShowJobDetailResponseStatus{
			value: "waiting",
		},
		RUNNING: ShowJobDetailResponseStatus{
			value: "running",
		},
	}
}

func (c ShowJobDetailResponseStatus) Value() string {
	return c.value
}

func (c ShowJobDetailResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobDetailResponseStatus) UnmarshalJSON(b []byte) error {
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
