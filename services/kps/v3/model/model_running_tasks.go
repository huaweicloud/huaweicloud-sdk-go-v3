package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 正在处理的任务详细信息。
type RunningTasks struct {

	// 虚拟机ID
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 操作类型。 - FAILED_RESET 重置 - FAILED_REPLACE 替换 - FAILED_UNBIND 解绑
	OperateType *RunningTasksOperateType `json:"operate_type,omitempty" xml:"operate_type"`

	// 任务时间
	TaskTime *string `json:"task_time,omitempty" xml:"task_time"`

	// 虚拟机名称
	ServerName *string `json:"server_name,omitempty" xml:"server_name"`

	// 虚拟机ID
	ServerId *string `json:"server_id,omitempty" xml:"server_id"`

	// 密钥对名称
	KeypairName *string `json:"keypair_name,omitempty" xml:"keypair_name"`
}

func (o RunningTasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunningTasks struct{}"
	}

	return strings.Join([]string{"RunningTasks", string(data)}, " ")
}

type RunningTasksOperateType struct {
	value string
}

type RunningTasksOperateTypeEnum struct {
	FAILED_RESET   RunningTasksOperateType
	FAILED_REPLACE RunningTasksOperateType
	FAILED_UNBIND  RunningTasksOperateType
}

func GetRunningTasksOperateTypeEnum() RunningTasksOperateTypeEnum {
	return RunningTasksOperateTypeEnum{
		FAILED_RESET: RunningTasksOperateType{
			value: "FAILED_RESET",
		},
		FAILED_REPLACE: RunningTasksOperateType{
			value: "FAILED_REPLACE",
		},
		FAILED_UNBIND: RunningTasksOperateType{
			value: "FAILED_UNBIND",
		},
	}
}

func (c RunningTasksOperateType) Value() string {
	return c.value
}

func (c RunningTasksOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunningTasksOperateType) UnmarshalJSON(b []byte) error {
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
