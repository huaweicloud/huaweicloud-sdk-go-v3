package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 任务详情
type DigitalHumanModelingJobInfo struct {

	// 照片建模任务ID。
	JobId string `json:"job_id"`

	// 任务的状态。 * WAITING：等待任务调度 * PROCESSING：正在处理 * PARTIAL_SUCCEED: 部分成功（模型生成，截图失败） * SUCCEED：成功 * FAILED：失败 * CANCELED：取消
	State DigitalHumanModelingJobInfoState `json:"state"`

	// 任务开始时间,格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间,格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	EndTime *string `json:"end_time,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`
}

func (o DigitalHumanModelingJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DigitalHumanModelingJobInfo struct{}"
	}

	return strings.Join([]string{"DigitalHumanModelingJobInfo", string(data)}, " ")
}

type DigitalHumanModelingJobInfoState struct {
	value string
}

type DigitalHumanModelingJobInfoStateEnum struct {
	WAITING         DigitalHumanModelingJobInfoState
	PROCESSING      DigitalHumanModelingJobInfoState
	PARTIAL_SUCCEED DigitalHumanModelingJobInfoState
	SUCCEED         DigitalHumanModelingJobInfoState
	FAILED          DigitalHumanModelingJobInfoState
	CANCELED        DigitalHumanModelingJobInfoState
}

func GetDigitalHumanModelingJobInfoStateEnum() DigitalHumanModelingJobInfoStateEnum {
	return DigitalHumanModelingJobInfoStateEnum{
		WAITING: DigitalHumanModelingJobInfoState{
			value: "WAITING",
		},
		PROCESSING: DigitalHumanModelingJobInfoState{
			value: "PROCESSING",
		},
		PARTIAL_SUCCEED: DigitalHumanModelingJobInfoState{
			value: "PARTIAL_SUCCEED",
		},
		SUCCEED: DigitalHumanModelingJobInfoState{
			value: "SUCCEED",
		},
		FAILED: DigitalHumanModelingJobInfoState{
			value: "FAILED",
		},
		CANCELED: DigitalHumanModelingJobInfoState{
			value: "CANCELED",
		},
	}
}

func (c DigitalHumanModelingJobInfoState) Value() string {
	return c.value
}

func (c DigitalHumanModelingJobInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalHumanModelingJobInfoState) UnmarshalJSON(b []byte) error {
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
