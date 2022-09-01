package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListAgentDimensionInfoRequest struct {

	// 发送的实体的MIME类型。默认使用application/json; charset=UTF-8。
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// 资源ID，如：4270ff17-aba3-4138-89fa-820594c39755。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 维度名称，枚举类型，类型有：   mount_point：挂载点,   disk：磁盘,   proc：进程,   gpu：显卡,   raid: RAID控制器,
	DimName ListAgentDimensionInfoRequestDimName `json:"dim_name" xml:"dim_name"`

	// 维度值，32位字符串，如：2e84018fc8b4484b94e89aae212fe615。
	DimValue *string `json:"dim_value,omitempty" xml:"dim_value"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
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
