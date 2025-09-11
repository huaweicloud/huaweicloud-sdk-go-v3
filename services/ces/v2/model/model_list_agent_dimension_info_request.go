package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAgentDimensionInfoRequest Request Object
type ListAgentDimensionInfoRequest struct {

	// **参数描述**： 资源ID，如：4270ff17-aba3-4138-89fa-820594c39755。 **约束限制**： 不涉及。 **取值范围**： 字符串长度为36。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数描述**： 维度名称。 **约束限制**： 不涉及。 **取值范围**： 枚举类型，mount_point:挂载点，disk:磁盘，proc:进程，gpu:显卡，raid:RAID控制器。 **默认取值**： 不涉及。
	DimName ListAgentDimensionInfoRequestDimName `json:"dim_name"`

	// **参数描述**： 维度值(建议：同一个instance_id下相同dim_value对应的原始维度值origin_value是一样的，无需多次调用，建议根据instance_id + dim_value作为key进行缓存重复使用)。 **约束限制**： 不涉及。  **取值范围**： 32位字符串，如：2e84018fc8b4484b94e89aae212fe615。 **默认取值**： 不涉及。
	DimValue *string `json:"dim_value,omitempty"`

	// **参数描述**： 分页偏移量。 **约束限制**： 不涉及。 **取值范围**： 最小值为0，最大值为2147483647。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数描述**： 分页大小。 **约束限制**： 不涉及。 **取值范围**： 最小值为1，最大值为1000。 **默认取值**： 1000
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAgentDimensionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentDimensionInfoRequest struct{}"
	}

	return strings.Join([]string{"ListAgentDimensionInfoRequest", string(data)}, " ")
}

type ListAgentDimensionInfoRequestDimName struct {
	value string
}

type ListAgentDimensionInfoRequestDimNameEnum struct {
	MOUNT_POINT ListAgentDimensionInfoRequestDimName
	DISK        ListAgentDimensionInfoRequestDimName
	PROC        ListAgentDimensionInfoRequestDimName
	GPU         ListAgentDimensionInfoRequestDimName
	RAID        ListAgentDimensionInfoRequestDimName
}

func GetListAgentDimensionInfoRequestDimNameEnum() ListAgentDimensionInfoRequestDimNameEnum {
	return ListAgentDimensionInfoRequestDimNameEnum{
		MOUNT_POINT: ListAgentDimensionInfoRequestDimName{
			value: "mount_point",
		},
		DISK: ListAgentDimensionInfoRequestDimName{
			value: "disk",
		},
		PROC: ListAgentDimensionInfoRequestDimName{
			value: "proc",
		},
		GPU: ListAgentDimensionInfoRequestDimName{
			value: "gpu",
		},
		RAID: ListAgentDimensionInfoRequestDimName{
			value: "raid",
		},
	}
}

func (c ListAgentDimensionInfoRequestDimName) Value() string {
	return c.value
}

func (c ListAgentDimensionInfoRequestDimName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentDimensionInfoRequestDimName) UnmarshalJSON(b []byte) error {
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
