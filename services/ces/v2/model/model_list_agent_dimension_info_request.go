package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAgentDimensionInfoRequest Request Object
type ListAgentDimensionInfoRequest struct {

	// 资源ID，如：4270ff17-aba3-4138-89fa-820594c39755。
	InstanceId string `json:"instance_id"`

	// 维度名称，枚举类型，类型有：   mount_point：挂载点,   disk：磁盘,   proc：进程,   gpu：显卡,   raid: RAID控制器,
	DimName ListAgentDimensionInfoRequestDimName `json:"dim_name"`

	// 维度值，32位字符串，如：2e84018fc8b4484b94e89aae212fe615。
	DimValue *string `json:"dim_value,omitempty"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
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
