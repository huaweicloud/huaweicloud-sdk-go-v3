package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PartitionConsumingStates struct {

	// 通道的分区标识符。 可定义为如下两种样式： - shardId-0000000000 - 0 比如一个通道有三个分区，那么分区标识符分别为0, 1, 2，或者shardId-0000000000, shardId-0000000001, shardId-0000000002
	PartitionId *string `json:"partition_id,omitempty"`

	// 需要提交的序列号，用来记录该通道的消费检查点，需要保证该序列号处于有效范围内。
	SequenceNumber *string `json:"sequence_number,omitempty"`

	// 索引位置, 最新的一条索引位置。
	LatestOffset *int64 `json:"latest_offset,omitempty"`

	// 索引位置, 最早的一条索引位置。
	EarliestOffset *int64 `json:"earliest_offset,omitempty"`

	// Checkpoint类型。 - LAST_READ：在数据库中只记录序列号。
	CheckpointType *PartitionConsumingStatesCheckpointType `json:"checkpoint_type,omitempty"`

	// 分区的当前状态。 - CREATING：创建中 - ACTIVE：可用 - DELETED：删除中 - EXPIRED：已过期
	Status *PartitionConsumingStatesStatus `json:"status,omitempty"`
}

func (o PartitionConsumingStates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionConsumingStates struct{}"
	}

	return strings.Join([]string{"PartitionConsumingStates", string(data)}, " ")
}

type PartitionConsumingStatesCheckpointType struct {
	value string
}

type PartitionConsumingStatesCheckpointTypeEnum struct {
	LAST_READ PartitionConsumingStatesCheckpointType
}

func GetPartitionConsumingStatesCheckpointTypeEnum() PartitionConsumingStatesCheckpointTypeEnum {
	return PartitionConsumingStatesCheckpointTypeEnum{
		LAST_READ: PartitionConsumingStatesCheckpointType{
			value: "LAST_READ",
		},
	}
}

func (c PartitionConsumingStatesCheckpointType) Value() string {
	return c.value
}

func (c PartitionConsumingStatesCheckpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PartitionConsumingStatesCheckpointType) UnmarshalJSON(b []byte) error {
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

type PartitionConsumingStatesStatus struct {
	value string
}

type PartitionConsumingStatesStatusEnum struct {
	CREATING PartitionConsumingStatesStatus
	ACTIVE   PartitionConsumingStatesStatus
	DELETED  PartitionConsumingStatesStatus
	EXPIRED  PartitionConsumingStatesStatus
}

func GetPartitionConsumingStatesStatusEnum() PartitionConsumingStatesStatusEnum {
	return PartitionConsumingStatesStatusEnum{
		CREATING: PartitionConsumingStatesStatus{
			value: "CREATING",
		},
		ACTIVE: PartitionConsumingStatesStatus{
			value: "ACTIVE",
		},
		DELETED: PartitionConsumingStatesStatus{
			value: "DELETED",
		},
		EXPIRED: PartitionConsumingStatesStatus{
			value: "EXPIRED",
		},
	}
}

func (c PartitionConsumingStatesStatus) Value() string {
	return c.value
}

func (c PartitionConsumingStatesStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PartitionConsumingStatesStatus) UnmarshalJSON(b []byte) error {
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
