package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowTaskResponse Response Object
type ShowTaskResponse struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 用户输入
	InputJson *interface{} `json:"input_json,omitempty"`

	// 状态
	Status *ShowTaskResponseStatus `json:"status,omitempty"`

	// 任务进度
	Progress *interface{} `json:"progress,omitempty"`

	// 输出
	OutputJson *interface{} `json:"output_json,omitempty"`

	// 开始时间
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 结束时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 创建时间
	CreateTime     *sdktime.SdkTime `json:"create_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskResponse", string(data)}, " ")
}

type ShowTaskResponseStatus struct {
	value string
}

type ShowTaskResponseStatusEnum struct {
	RUNNING   ShowTaskResponseStatus
	PENDING   ShowTaskResponseStatus
	EXPIRED   ShowTaskResponseStatus
	UNKNOWN   ShowTaskResponseStatus
	FAILED    ShowTaskResponseStatus
	SUCCEEDED ShowTaskResponseStatus
	STOPPED   ShowTaskResponseStatus
	DELETED   ShowTaskResponseStatus
}

func GetShowTaskResponseStatusEnum() ShowTaskResponseStatusEnum {
	return ShowTaskResponseStatusEnum{
		RUNNING: ShowTaskResponseStatus{
			value: "Running",
		},
		PENDING: ShowTaskResponseStatus{
			value: "Pending",
		},
		EXPIRED: ShowTaskResponseStatus{
			value: "Expired",
		},
		UNKNOWN: ShowTaskResponseStatus{
			value: "Unknown",
		},
		FAILED: ShowTaskResponseStatus{
			value: "Failed",
		},
		SUCCEEDED: ShowTaskResponseStatus{
			value: "Succeeded",
		},
		STOPPED: ShowTaskResponseStatus{
			value: "Stopped",
		},
		DELETED: ShowTaskResponseStatus{
			value: "Deleted",
		},
	}
}

func (c ShowTaskResponseStatus) Value() string {
	return c.value
}

func (c ShowTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
