package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AgentDimension struct {

	// **参数解释**： 维度名称。 **取值范围**： 枚举类型，mount_point:挂载点，disk:磁盘，proc:进程，gpu:显卡，raid:RAID控制器。
	Name *AgentDimensionName `json:"name,omitempty"`

	// **参数解释**： 维度值。 **取值范围**： 字符串长度为32。
	Value *string `json:"value,omitempty"`

	// **参数解释**： 实际维度信息。 **取值范围** 字符串长度在 1 到 1024 之间。
	OriginValue *string `json:"origin_value,omitempty"`
}

func (o AgentDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentDimension struct{}"
	}

	return strings.Join([]string{"AgentDimension", string(data)}, " ")
}

type AgentDimensionName struct {
	value string
}

type AgentDimensionNameEnum struct {
	MOUNT_POINT AgentDimensionName
	DISK        AgentDimensionName
	PROC        AgentDimensionName
	GPU         AgentDimensionName
	RAID        AgentDimensionName
}

func GetAgentDimensionNameEnum() AgentDimensionNameEnum {
	return AgentDimensionNameEnum{
		MOUNT_POINT: AgentDimensionName{
			value: "mount_point",
		},
		DISK: AgentDimensionName{
			value: "disk",
		},
		PROC: AgentDimensionName{
			value: "proc",
		},
		GPU: AgentDimensionName{
			value: "gpu",
		},
		RAID: AgentDimensionName{
			value: "raid",
		},
	}
}

func (c AgentDimensionName) Value() string {
	return c.value
}

func (c AgentDimensionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgentDimensionName) UnmarshalJSON(b []byte) error {
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
