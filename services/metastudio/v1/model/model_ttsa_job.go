package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TtsaJob struct {

	// 语音驱动任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务的状态。 * WAITING * PROCESSING * SUCCEED * FAILED
	State *TtsaJobState `json:"state,omitempty"`

	// 任务开始时间,格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间,格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	EndTime *string `json:"end_time,omitempty"`

	// 语音驱动内容时长，单位:秒。
	ContentDuration *float32 `json:"content_duration,omitempty"`
}

func (o TtsaJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtsaJob struct{}"
	}

	return strings.Join([]string{"TtsaJob", string(data)}, " ")
}

type TtsaJobState struct {
	value string
}

type TtsaJobStateEnum struct {
	WAITING    TtsaJobState
	PROCESSING TtsaJobState
	SUCCEED    TtsaJobState
	FAILED     TtsaJobState
}

func GetTtsaJobStateEnum() TtsaJobStateEnum {
	return TtsaJobStateEnum{
		WAITING: TtsaJobState{
			value: "WAITING",
		},
		PROCESSING: TtsaJobState{
			value: "PROCESSING",
		},
		SUCCEED: TtsaJobState{
			value: "SUCCEED",
		},
		FAILED: TtsaJobState{
			value: "FAILED",
		},
	}
}

func (c TtsaJobState) Value() string {
	return c.value
}

func (c TtsaJobState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TtsaJobState) UnmarshalJSON(b []byte) error {
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
