package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowConsumerStateRequest Request Object
type ShowConsumerStateRequest struct {

	// 需要查询的App名称。
	AppName string `json:"app_name"`

	// 需要查询的通道名称。
	StreamName string `json:"stream_name"`

	// 单次请求返回的最大分区数。最小值是1，最大值是1000；默认值是100。
	Limit *int32 `json:"limit,omitempty"`

	// 从该分区值开始返回分区列表，返回的分区列表不包括此分区。
	StartPartitionId *string `json:"start_partition_id,omitempty"`

	// Checkpoint类型。  - LAST_READ：在数据库中只记录序列号。
	CheckpointType ShowConsumerStateRequestCheckpointType `json:"checkpoint_type"`
}

func (o ShowConsumerStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerStateRequest struct{}"
	}

	return strings.Join([]string{"ShowConsumerStateRequest", string(data)}, " ")
}

type ShowConsumerStateRequestCheckpointType struct {
	value string
}

type ShowConsumerStateRequestCheckpointTypeEnum struct {
	LAST_READ ShowConsumerStateRequestCheckpointType
}

func GetShowConsumerStateRequestCheckpointTypeEnum() ShowConsumerStateRequestCheckpointTypeEnum {
	return ShowConsumerStateRequestCheckpointTypeEnum{
		LAST_READ: ShowConsumerStateRequestCheckpointType{
			value: "LAST_READ",
		},
	}
}

func (c ShowConsumerStateRequestCheckpointType) Value() string {
	return c.value
}

func (c ShowConsumerStateRequestCheckpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConsumerStateRequestCheckpointType) UnmarshalJSON(b []byte) error {
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
