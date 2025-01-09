package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteDesktopPoolScriptsReq 桌面池执行脚本请求体。
type ExecuteDesktopPoolScriptsReq struct {

	// 执行的脚本id列表，与command二选一。
	ScriptIds *[]string `json:"script_ids,omitempty"`

	// 首批执行的桌面数量，优先级高于gray_desktop_ids。
	GrayCount *int32 `json:"gray_count,omitempty"`

	// 首批执行的桌面id列表，优先级低于gray_count。
	GrayDesktopIds *[]string `json:"gray_desktop_ids,omitempty"`

	// 灰度失败阈值，灰度执行失败次数达到该值时，不执行下一批任务。
	GrayFailThreshold *int32 `json:"gray_fail_threshold,omitempty"`

	// 执行脚本前置步骤。
	PreStart *string `json:"pre_start,omitempty"`

	// 执行脚本完成后置步骤，当前支持关机（stop）、重启（reboot）。
	PostFinish *string `json:"post_finish,omitempty"`

	// 执行的命令行，与script_ids二选一。
	CommandContent *string `json:"command_content,omitempty"`

	// 命令行类型，执行命令行时必传。 - POWERSHELL：WINDOWS系统使用。 - BAT：WINDOWS系统使用。 - SHELL：LINUX系统使用。
	CommandType *ExecuteDesktopPoolScriptsReqCommandType `json:"command_type,omitempty"`

	// 执行脚本的超时时间，单位分钟，默认1分钟。
	ExecutionTimeout *int32 `json:"execution_timeout,omitempty"`
}

func (o ExecuteDesktopPoolScriptsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDesktopPoolScriptsReq struct{}"
	}

	return strings.Join([]string{"ExecuteDesktopPoolScriptsReq", string(data)}, " ")
}

type ExecuteDesktopPoolScriptsReqCommandType struct {
	value string
}

type ExecuteDesktopPoolScriptsReqCommandTypeEnum struct {
	POWERSHELL ExecuteDesktopPoolScriptsReqCommandType
	BAT        ExecuteDesktopPoolScriptsReqCommandType
	SHELL      ExecuteDesktopPoolScriptsReqCommandType
}

func GetExecuteDesktopPoolScriptsReqCommandTypeEnum() ExecuteDesktopPoolScriptsReqCommandTypeEnum {
	return ExecuteDesktopPoolScriptsReqCommandTypeEnum{
		POWERSHELL: ExecuteDesktopPoolScriptsReqCommandType{
			value: "POWERSHELL",
		},
		BAT: ExecuteDesktopPoolScriptsReqCommandType{
			value: "BAT",
		},
		SHELL: ExecuteDesktopPoolScriptsReqCommandType{
			value: "SHELL",
		},
	}
}

func (c ExecuteDesktopPoolScriptsReqCommandType) Value() string {
	return c.value
}

func (c ExecuteDesktopPoolScriptsReqCommandType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteDesktopPoolScriptsReqCommandType) UnmarshalJSON(b []byte) error {
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
