package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScriptExecutionTask 脚本执行任务信息。
type ScriptExecutionTask struct {

	// 任务id。
	Id *string `json:"id,omitempty"`

	// 桌面池id。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 桌面池名称。
	DesktopPoolName *string `json:"desktop_pool_name,omitempty"`

	// 脚本信息列表。
	Scripts *[]Script `json:"scripts,omitempty"`

	// 执行的命令行。
	CommandContent *string `json:"command_content,omitempty"`

	// 命令行类型。 - POWERSHELL：WINDOWS系统使用。 - BAT：WINDOWS系统使用。 - SHELL：LINUX系统使用。
	CommandType *string `json:"command_type,omitempty"`

	// 任务开始时间，格式为：yyyy-MM-ddTHH:mm:ssZ。
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间，格式为：yyyy-MM-ddTHH:mm:ssZ。
	EndTime *string `json:"end_time,omitempty"`

	// 任务状态，值含： - FINISH：已完成。 - FAILED：失败。 - RUNNING：运行中。 - INIT： 初始化。
	Status *ScriptExecutionTaskStatus `json:"status,omitempty"`

	// 成功数量。
	SuccessNum *int32 `json:"success_num,omitempty"`

	// 失败数量。
	FailedNum *int32 `json:"failed_num,omitempty"`

	// 跳过数量。
	SkipNum *int32 `json:"skip_num,omitempty"`
}

func (o ScriptExecutionTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptExecutionTask struct{}"
	}

	return strings.Join([]string{"ScriptExecutionTask", string(data)}, " ")
}

type ScriptExecutionTaskStatus struct {
	value string
}

type ScriptExecutionTaskStatusEnum struct {
	FINISH  ScriptExecutionTaskStatus
	FAILED  ScriptExecutionTaskStatus
	RUNNING ScriptExecutionTaskStatus
	INIT    ScriptExecutionTaskStatus
}

func GetScriptExecutionTaskStatusEnum() ScriptExecutionTaskStatusEnum {
	return ScriptExecutionTaskStatusEnum{
		FINISH: ScriptExecutionTaskStatus{
			value: "FINISH",
		},
		FAILED: ScriptExecutionTaskStatus{
			value: "FAILED",
		},
		RUNNING: ScriptExecutionTaskStatus{
			value: "RUNNING",
		},
		INIT: ScriptExecutionTaskStatus{
			value: "INIT",
		},
	}
}

func (c ScriptExecutionTaskStatus) Value() string {
	return c.value
}

func (c ScriptExecutionTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptExecutionTaskStatus) UnmarshalJSON(b []byte) error {
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
