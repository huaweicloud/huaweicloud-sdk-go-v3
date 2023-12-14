package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowCheckpointRequest Request Object
type ShowCheckpointRequest struct {

	// 该Checkpoint所属的通道名称。
	StreamName string `json:"stream_name"`

	// 该Checkpoint所属的通道分区标识符。  可定义为如下两种样式：  - shardId-0000000000 - 0  比如一个通道有三个分区，那么分区标识符分别为0, 1, 2，或者shardId-0000000000, shardId-0000000001, shardId-0000000002
	PartitionId string `json:"partition_id"`

	// 该Checkpoint关联App名称。
	AppName string `json:"app_name"`

	// Checkpoint类型。  - LAST_READ：在数据库中只记录序列号。
	CheckpointType ShowCheckpointRequestCheckpointType `json:"checkpoint_type"`
}

func (o ShowCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckpointRequest struct{}"
	}

	return strings.Join([]string{"ShowCheckpointRequest", string(data)}, " ")
}

type ShowCheckpointRequestCheckpointType struct {
	value string
}

type ShowCheckpointRequestCheckpointTypeEnum struct {
	LAST_READ ShowCheckpointRequestCheckpointType
}

func GetShowCheckpointRequestCheckpointTypeEnum() ShowCheckpointRequestCheckpointTypeEnum {
	return ShowCheckpointRequestCheckpointTypeEnum{
		LAST_READ: ShowCheckpointRequestCheckpointType{
			value: "LAST_READ",
		},
	}
}

func (c ShowCheckpointRequestCheckpointType) Value() string {
	return c.value
}

func (c ShowCheckpointRequestCheckpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCheckpointRequestCheckpointType) UnmarshalJSON(b []byte) error {
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
