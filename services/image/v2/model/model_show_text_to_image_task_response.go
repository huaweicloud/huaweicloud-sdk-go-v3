package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowTextToImageTaskResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务创建时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	CreateTime *string `json:"create_time,omitempty"`

	// 任务更新时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	UpdateTime *string `json:"update_time,omitempty"`

	// 任务当前的状态，分别为SUCCEEDED（运行成功），FAILED（运行失败），RUNNING（运行中）。
	State *ShowTextToImageTaskResponseState `json:"state,omitempty"`

	Output *TaskOutput `json:"output,omitempty"`

	Config *TextToImageTaskConfig `json:"config,omitempty"`

	Callback *TaskCallback `json:"callback,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTextToImageTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTextToImageTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTextToImageTaskResponse", string(data)}, " ")
}

type ShowTextToImageTaskResponseState struct {
	value string
}

type ShowTextToImageTaskResponseStateEnum struct {
	SUCCEEDED ShowTextToImageTaskResponseState
	FAILED    ShowTextToImageTaskResponseState
	RUNNING   ShowTextToImageTaskResponseState
}

func GetShowTextToImageTaskResponseStateEnum() ShowTextToImageTaskResponseStateEnum {
	return ShowTextToImageTaskResponseStateEnum{
		SUCCEEDED: ShowTextToImageTaskResponseState{
			value: "SUCCEEDED",
		},
		FAILED: ShowTextToImageTaskResponseState{
			value: "FAILED",
		},
		RUNNING: ShowTextToImageTaskResponseState{
			value: "RUNNING",
		},
	}
}

func (c ShowTextToImageTaskResponseState) Value() string {
	return c.value
}

func (c ShowTextToImageTaskResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTextToImageTaskResponseState) UnmarshalJSON(b []byte) error {
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
