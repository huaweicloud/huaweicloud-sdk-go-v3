package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExtensionInfo struct {

	// **参数解释**: 插件名称 **取值范围**: 长度为[1,64]个字符
	Name *string `json:"name,omitempty"`

	// **参数解释**:  插件状态 **取值范围**: - none: 未安装 - running: 运行中 - stopped: 已停止 - fault: 故障（进程异常） - unknown: 故障（连接异常）
	Status *ExtensionInfoStatus `json:"status,omitempty"`

	// **参数解释**: 插件版本 **取值范围**: 长度为[1,32]个字符
	Version *string `json:"version,omitempty"`
}

func (o ExtensionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionInfo struct{}"
	}

	return strings.Join([]string{"ExtensionInfo", string(data)}, " ")
}

type ExtensionInfoStatus struct {
	value string
}

type ExtensionInfoStatusEnum struct {
	NONE    ExtensionInfoStatus
	RUNNING ExtensionInfoStatus
	STOPPED ExtensionInfoStatus
	FAULT   ExtensionInfoStatus
	UNKNOWN ExtensionInfoStatus
}

func GetExtensionInfoStatusEnum() ExtensionInfoStatusEnum {
	return ExtensionInfoStatusEnum{
		NONE: ExtensionInfoStatus{
			value: "none",
		},
		RUNNING: ExtensionInfoStatus{
			value: "running",
		},
		STOPPED: ExtensionInfoStatus{
			value: "stopped",
		},
		FAULT: ExtensionInfoStatus{
			value: "fault",
		},
		UNKNOWN: ExtensionInfoStatus{
			value: "unknown",
		},
	}
}

func (c ExtensionInfoStatus) Value() string {
	return c.value
}

func (c ExtensionInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtensionInfoStatus) UnmarshalJSON(b []byte) error {
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
