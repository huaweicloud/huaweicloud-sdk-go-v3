package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type RunQueryVideoModerationJobResponse struct {

	// 本次请求的唯⼀标识，⽤于问题排查，建议保存。
	RequestId *string `json:"request_id,omitempty"`

	// 作业id
	JobId *string `json:"job_id,omitempty"`

	// 作业状态，可取值有：  running: 正在运行 succeeded: 运行成功  failed: 运行失败
	Status *RunQueryVideoModerationJobResponseStatus `json:"status,omitempty"`

	RequestParams *VideoModerationResultRequestParams `json:"request_params,omitempty"`

	// 作业创建时间
	CraeteTime *string `json:"craete_time,omitempty"`

	// 作业更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	Result         *VideoModerationResultResult `json:"result,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o RunQueryVideoModerationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryVideoModerationJobResponse struct{}"
	}

	return strings.Join([]string{"RunQueryVideoModerationJobResponse", string(data)}, " ")
}

type RunQueryVideoModerationJobResponseStatus struct {
	value string
}

type RunQueryVideoModerationJobResponseStatusEnum struct {
	RUNNING   RunQueryVideoModerationJobResponseStatus
	SUCCEEDED RunQueryVideoModerationJobResponseStatus
	FAILED    RunQueryVideoModerationJobResponseStatus
}

func GetRunQueryVideoModerationJobResponseStatusEnum() RunQueryVideoModerationJobResponseStatusEnum {
	return RunQueryVideoModerationJobResponseStatusEnum{
		RUNNING: RunQueryVideoModerationJobResponseStatus{
			value: "running",
		},
		SUCCEEDED: RunQueryVideoModerationJobResponseStatus{
			value: "succeeded",
		},
		FAILED: RunQueryVideoModerationJobResponseStatus{
			value: "failed",
		},
	}
}

func (c RunQueryVideoModerationJobResponseStatus) Value() string {
	return c.value
}

func (c RunQueryVideoModerationJobResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunQueryVideoModerationJobResponseStatus) UnmarshalJSON(b []byte) error {
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
