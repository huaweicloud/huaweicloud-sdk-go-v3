package model

import (
	"encoding/json"
	"os"

	"strings"
)

type PutRecordsRequestEntry struct {
	// 需要上传的数据。  上传的数据为序列化之后的二进制数据（Base64编码后的字符串）。  比如需要上传字符串“data”，“data”经过Base64编码之后是“ZGF0YQ==”。
	Data **os.File `json:"data"`
	// 用于明确数据需要写入分区的哈希值，此哈希值将覆盖“partition_key”的哈希值。  取值范围：0~long.max
	ExplicitHashKey *string `json:"explicit_hash_key,omitempty"`
	// 通道的分区标识符。 可定义为如下两种样式： - shardId-0000000000 - 0  比如一个通道有三个分区，那么分区标识符分别为0, 1, 2，或者shardId-0000000000, shardId-0000000001, shardId-0000000002
	PartitionId *string `json:"partition_id,omitempty"`
	// 数据将写入的分区。  说明：  如果传了partition_id参数，则优先使用partition_id参数。如果partition_id没有传，则使用partition_key。
	PartitionKey *string `json:"partition_key,omitempty"`
	Timestamp    *int64  `json:"timestamp,omitempty"`
}

func (o PutRecordsRequestEntry) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PutRecordsRequestEntry struct{}"
	}

	return strings.Join([]string{"PutRecordsRequestEntry", string(data)}, " ")
}
