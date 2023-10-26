package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OneFsTaskResp 查询任务详情
type OneFsTaskResp struct {

	// 任务id
	TaskId string `json:"task_id"`

	// 任务状态, SUCCESS表示成功，DOING表示正在执行，FAIL表示失败
	Status OneFsTaskRespStatus `json:"status"`

	// 目录资源使用情况(包含子目录)
	DirUsage *FsDuInfo `json:"dir_usage,omitempty"`

	// 任务开始时间，UTC时间，例如：2006-01-02 15:04:05'
	BeginTime string `json:"begin_time"`

	// 任务结束时间，UTC时间，例如：2006-01-02 15:04:06'
	EndTime string `json:"end_time"`
}

func (o OneFsTaskResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneFsTaskResp struct{}"
	}

	return strings.Join([]string{"OneFsTaskResp", string(data)}, " ")
}

type OneFsTaskRespStatus struct {
	value string
}

type OneFsTaskRespStatusEnum struct {
	SUCCESS OneFsTaskRespStatus
	DOING   OneFsTaskRespStatus
	FAIL    OneFsTaskRespStatus
}

func GetOneFsTaskRespStatusEnum() OneFsTaskRespStatusEnum {
	return OneFsTaskRespStatusEnum{
		SUCCESS: OneFsTaskRespStatus{
			value: "SUCCESS",
		},
		DOING: OneFsTaskRespStatus{
			value: "DOING",
		},
		FAIL: OneFsTaskRespStatus{
			value: "FAIL",
		},
	}
}

func (c OneFsTaskRespStatus) Value() string {
	return c.value
}

func (c OneFsTaskRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OneFsTaskRespStatus) UnmarshalJSON(b []byte) error {
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
