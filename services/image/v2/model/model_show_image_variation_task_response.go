package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowImageVariationTaskResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务创建时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	CreateTime *string `json:"create_time,omitempty"`

	// 任务更新时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	UpdateTime *string `json:"update_time,omitempty"`

	// 任务当前的状态，分别为SUCCEEDED（运行成功），FAILED（运行失败），RUNNING（运行中）。
	State *ShowImageVariationTaskResponseState `json:"state,omitempty"`

	Input *TaskInput `json:"input,omitempty"`

	Output *TaskOutput `json:"output,omitempty"`

	Config *ImageVariationTaskConfig `json:"config,omitempty"`

	Callback *TaskCallback `json:"callback,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowImageVariationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageVariationTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowImageVariationTaskResponse", string(data)}, " ")
}

type ShowImageVariationTaskResponseState struct {
	value string
}

type ShowImageVariationTaskResponseStateEnum struct {
	SUCCEEDED ShowImageVariationTaskResponseState
	FAILED    ShowImageVariationTaskResponseState
	RUNNING   ShowImageVariationTaskResponseState
}

func GetShowImageVariationTaskResponseStateEnum() ShowImageVariationTaskResponseStateEnum {
	return ShowImageVariationTaskResponseStateEnum{
		SUCCEEDED: ShowImageVariationTaskResponseState{
			value: "SUCCEEDED",
		},
		FAILED: ShowImageVariationTaskResponseState{
			value: "FAILED",
		},
		RUNNING: ShowImageVariationTaskResponseState{
			value: "RUNNING",
		},
	}
}

func (c ShowImageVariationTaskResponseState) Value() string {
	return c.value
}

func (c ShowImageVariationTaskResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageVariationTaskResponseState) UnmarshalJSON(b []byte) error {
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
