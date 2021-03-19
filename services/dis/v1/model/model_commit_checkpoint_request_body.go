package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CommitCheckpointRequestBody struct {
	// APP的名称，用户数据消费程序的唯一标识符，需要先通过创建App接口创建。

	AppName string `json:"app_name"`
	// Checkpoint类型。  - LAST_READ：在数据库中只记录序列号。

	CheckpointType CommitCheckpointRequestBodyCheckpointType `json:"checkpoint_type"`
	// 已创建的通道名称。

	StreamName string `json:"stream_name"`
	// 通道的分区标识符。 可定义为如下两种样式： - shardId-0000000000 - 0 比如一个通道有三个分区，那么分区标识符分别为0, 1, 2，或者shardId-0000000000, shardId-0000000001, shardId-0000000002

	PartitionId string `json:"partition_id"`
	// 需要提交的序列号，用来记录该通道的消费检查点，需要保证该序列号处于有效范围内。

	SequenceNumber string `json:"sequence_number"`
	// 用户消费程序端的元数据信息。  元数据信息的最大长度为1000个字符。

	Metadata *string `json:"metadata,omitempty"`
}

func (o CommitCheckpointRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CommitCheckpointRequestBody struct{}"
	}

	return strings.Join([]string{"CommitCheckpointRequestBody", string(data)}, " ")
}

type CommitCheckpointRequestBodyCheckpointType struct {
	value string
}

type CommitCheckpointRequestBodyCheckpointTypeEnum struct {
	LAST_READ CommitCheckpointRequestBodyCheckpointType
}

func GetCommitCheckpointRequestBodyCheckpointTypeEnum() CommitCheckpointRequestBodyCheckpointTypeEnum {
	return CommitCheckpointRequestBodyCheckpointTypeEnum{
		LAST_READ: CommitCheckpointRequestBodyCheckpointType{
			value: "LAST_READ",
		},
	}
}

func (c CommitCheckpointRequestBodyCheckpointType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CommitCheckpointRequestBodyCheckpointType) UnmarshalJSON(b []byte) error {
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
