package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CommitCheckpointReq struct {

	// APP的名称，用户数据消费程序的唯一标识符，需要先通过创建App接口创建。
	AppName string `json:"app_name"`

	// Checkpoint类型。  - LAST_READ：在数据库中只记录序列号。
	CheckpointType CommitCheckpointReqCheckpointType `json:"checkpoint_type"`

	// 已创建的通道名称。
	StreamName string `json:"stream_name"`

	// 通道的分区标识符。 可定义为如下两种样式： - shardId-0000000000 - 0 比如一个通道有三个分区，那么分区标识符分别为0, 1, 2，或者shardId-0000000000, shardId-0000000001, shardId-0000000002
	PartitionId string `json:"partition_id"`

	// 需要提交的序列号，用来记录该通道的消费检查点，需要保证该序列号处于有效范围内。
	SequenceNumber string `json:"sequence_number"`

	// 用户消费程序端的元数据信息。  元数据信息的最大长度为1000个字符。
	Metadata *string `json:"metadata,omitempty"`
}

func (o CommitCheckpointReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitCheckpointReq struct{}"
	}

	return strings.Join([]string{"CommitCheckpointReq", string(data)}, " ")
}

type CommitCheckpointReqCheckpointType struct {
	value string
}

type CommitCheckpointReqCheckpointTypeEnum struct {
	LAST_READ CommitCheckpointReqCheckpointType
}

func GetCommitCheckpointReqCheckpointTypeEnum() CommitCheckpointReqCheckpointTypeEnum {
	return CommitCheckpointReqCheckpointTypeEnum{
		LAST_READ: CommitCheckpointReqCheckpointType{
			value: "LAST_READ",
		},
	}
}

func (c CommitCheckpointReqCheckpointType) Value() string {
	return c.value
}

func (c CommitCheckpointReqCheckpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommitCheckpointReqCheckpointType) UnmarshalJSON(b []byte) error {
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
