package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AgentDimension struct {

	// 维度名称，枚举类型，类型有：   mount_point：挂载点，   disk：磁盘，   proc：进程，   gpu：显卡，   raid: RAID控制器
	Name *AgentDimensionName `json:"name,omitempty" xml:"name"`

	// 维度值，32位字符串，如：2e84018fc8b4484b94e89aae212fe615
	Value *string `json:"value,omitempty" xml:"value"`

	// 实际维度信息，字符串，如：vda。
	OriginValue *string `json:"origin_value,omitempty" xml:"origin_value"`
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
