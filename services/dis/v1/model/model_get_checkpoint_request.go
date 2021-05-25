package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type GetCheckpointRequest struct {
	// 该Checkpoint所属的通道名称。

	StreamName string `json:"stream_name"`
	// 该Checkpoint所属的通道分区标识符。  可定义为如下两种样式：  - shardId-0000000000 - 0  比如一个通道有三个分区，那么分区标识符分别为0, 1, 2，或者shardId-0000000000, shardId-0000000001, shardId-0000000002

	PartitionId string `json:"partition_id"`
	// 该Checkpoint关联App名称。

	AppName string `json:"app_name"`
	// Checkpoint类型。  - LAST_READ：在数据库中只记录序列号。

	CheckpointType GetCheckpointRequestCheckpointType `json:"checkpoint_type"`
}

func (o GetCheckpointRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetCheckpointRequest struct{}"
	}

	return strings.Join([]string{"GetCheckpointRequest", string(data)}, " ")
}

type GetCheckpointRequestCheckpointType struct {
	value string
}

type GetCheckpointRequestCheckpointTypeEnum struct {
	LAST_READ GetCheckpointRequestCheckpointType
}

func GetGetCheckpointRequestCheckpointTypeEnum() GetCheckpointRequestCheckpointTypeEnum {
	return GetCheckpointRequestCheckpointTypeEnum{
		LAST_READ: GetCheckpointRequestCheckpointType{
			value: "LAST_READ",
		},
	}
}

func (c GetCheckpointRequestCheckpointType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GetCheckpointRequestCheckpointType) UnmarshalJSON(b []byte) error {
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
