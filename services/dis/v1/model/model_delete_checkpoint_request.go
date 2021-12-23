package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteCheckpointRequest struct {
	// 该Checkpoint所属的通道名称。

	StreamName string `json:"stream_name"`
	// 该Checkpoint关联App名称。

	AppName string `json:"app_name"`
	// Checkpoint类型。  LAST_READ：在数据库中只记录序列号。

	CheckpointType DeleteCheckpointRequestCheckpointType `json:"checkpoint_type"`
	// 该Checkpoint所属的通道分区标识符。  可定义为如下两种样式：  - shardId-0000000000 - 0  比如一个通道有三个分区，那么分区标识符分别为0, 1, 2，或者shardId-0000000000, shardId-0000000001, shardId-0000000002

	PartitionId *string `json:"partition_id,omitempty"`
}

func (o DeleteCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCheckpointRequest struct{}"
	}

	return strings.Join([]string{"DeleteCheckpointRequest", string(data)}, " ")
}

type DeleteCheckpointRequestCheckpointType struct {
	value string
}

type DeleteCheckpointRequestCheckpointTypeEnum struct {
	LAST_READ DeleteCheckpointRequestCheckpointType
}

func GetDeleteCheckpointRequestCheckpointTypeEnum() DeleteCheckpointRequestCheckpointTypeEnum {
	return DeleteCheckpointRequestCheckpointTypeEnum{
		LAST_READ: DeleteCheckpointRequestCheckpointType{
			value: "LAST_READ",
		},
	}
}

func (c DeleteCheckpointRequestCheckpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteCheckpointRequestCheckpointType) UnmarshalJSON(b []byte) error {
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
