package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowSqlLimitJobInfoResponse struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 任务状态
	JobStatus *ShowSqlLimitJobInfoResponseJobStatus `json:"job_status,omitempty"`

	// 失败原因
	FailReason     *string `json:"fail_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSqlLimitJobInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitJobInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitJobInfoResponse", string(data)}, " ")
}

type ShowSqlLimitJobInfoResponseJobStatus struct {
	value string
}

type ShowSqlLimitJobInfoResponseJobStatusEnum struct {
	RUNNING   ShowSqlLimitJobInfoResponseJobStatus
	COMPLETED ShowSqlLimitJobInfoResponseJobStatus
	FAILED    ShowSqlLimitJobInfoResponseJobStatus
}

func GetShowSqlLimitJobInfoResponseJobStatusEnum() ShowSqlLimitJobInfoResponseJobStatusEnum {
	return ShowSqlLimitJobInfoResponseJobStatusEnum{
		RUNNING: ShowSqlLimitJobInfoResponseJobStatus{
			value: "RUNNING",
		},
		COMPLETED: ShowSqlLimitJobInfoResponseJobStatus{
			value: "COMPLETED",
		},
		FAILED: ShowSqlLimitJobInfoResponseJobStatus{
			value: "FAILED",
		},
	}
}

func (c ShowSqlLimitJobInfoResponseJobStatus) Value() string {
	return c.value
}

func (c ShowSqlLimitJobInfoResponseJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlLimitJobInfoResponseJobStatus) UnmarshalJSON(b []byte) error {
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
