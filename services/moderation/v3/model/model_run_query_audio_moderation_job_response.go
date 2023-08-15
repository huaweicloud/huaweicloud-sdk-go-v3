package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RunQueryAudioModerationJobResponse Response Object
type RunQueryAudioModerationJobResponse struct {

	// 作业id
	JobId *string `json:"job_id,omitempty"`

	// 作业状态，可取值有： running: 正在运行 succeeded: 运行成功 failed: 运行失败
	Status *RunQueryAudioModerationJobResponseStatus `json:"status,omitempty"`

	Result *AudioModerationResultResult `json:"result,omitempty"`

	RequestParams *AudioModerationResultRequestParams `json:"request_params,omitempty"`

	// 作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 作业更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 本次请求的唯⼀标识，⽤于问题排查，建议保存。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunQueryAudioModerationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryAudioModerationJobResponse struct{}"
	}

	return strings.Join([]string{"RunQueryAudioModerationJobResponse", string(data)}, " ")
}

type RunQueryAudioModerationJobResponseStatus struct {
	value string
}

type RunQueryAudioModerationJobResponseStatusEnum struct {
	RUNNING   RunQueryAudioModerationJobResponseStatus
	SUCCEEDED RunQueryAudioModerationJobResponseStatus
	FAILED    RunQueryAudioModerationJobResponseStatus
}

func GetRunQueryAudioModerationJobResponseStatusEnum() RunQueryAudioModerationJobResponseStatusEnum {
	return RunQueryAudioModerationJobResponseStatusEnum{
		RUNNING: RunQueryAudioModerationJobResponseStatus{
			value: "running",
		},
		SUCCEEDED: RunQueryAudioModerationJobResponseStatus{
			value: "succeeded",
		},
		FAILED: RunQueryAudioModerationJobResponseStatus{
			value: "failed",
		},
	}
}

func (c RunQueryAudioModerationJobResponseStatus) Value() string {
	return c.value
}

func (c RunQueryAudioModerationJobResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunQueryAudioModerationJobResponseStatus) UnmarshalJSON(b []byte) error {
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
