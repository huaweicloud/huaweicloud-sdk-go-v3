package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFsTaskResponse Response Object
type ShowFsTaskResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务状态, SUCCESS表示成功，DOING表示正在执行，FAIL表示失败
	Status *ShowFsTaskResponseStatus `json:"status,omitempty"`

	// 目录资源使用情况(包含子目录)
	DirUsage *FsDuInfo `json:"dir_usage,omitempty"`

	// 任务开始时间，UTC时间，例如：2006-01-02 15:04:05'
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间，UTC时间，例如：2006-01-02 15:04:06'
	EndTime        *string `json:"end_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFsTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowFsTaskResponse", string(data)}, " ")
}

type ShowFsTaskResponseStatus struct {
	value string
}

type ShowFsTaskResponseStatusEnum struct {
	SUCCESS ShowFsTaskResponseStatus
	DOING   ShowFsTaskResponseStatus
	FAIL    ShowFsTaskResponseStatus
}

func GetShowFsTaskResponseStatusEnum() ShowFsTaskResponseStatusEnum {
	return ShowFsTaskResponseStatusEnum{
		SUCCESS: ShowFsTaskResponseStatus{
			value: "SUCCESS",
		},
		DOING: ShowFsTaskResponseStatus{
			value: "DOING",
		},
		FAIL: ShowFsTaskResponseStatus{
			value: "FAIL",
		},
	}
}

func (c ShowFsTaskResponseStatus) Value() string {
	return c.value
}

func (c ShowFsTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFsTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
