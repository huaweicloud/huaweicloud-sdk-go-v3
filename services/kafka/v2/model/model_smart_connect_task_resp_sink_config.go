package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartConnectTaskRespSinkConfig struct {

	// **参数解释**： 转储启动偏移量。（仅目标端类型为OBS时会显示） **取值范围**： - latest：获取最新的数据。 - earliest：获取最早的数据。
	ConsumerStrategy *string `json:"consumer_strategy,omitempty"`

	// **参数解释**： 转储文件格式。当前只支持TEXT。（仅目标端类型为OBS时会显示） **取值范围**： TEXT。
	DestinationFileType *string `json:"destination_file_type,omitempty"`

	// **参数解释**： 数据转储周期（秒）。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	DeliverTimeInterval *int32 `json:"deliver_time_interval,omitempty"`

	// **参数解释**： 转储地址。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	// **参数解释**： 转储目录。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	ObsPath *string `json:"obs_path,omitempty"`

	// **参数解释**： 时间目录格式。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	PartitionFormat *string `json:"partition_format,omitempty"`

	// **参数解释**： 记录分行符。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	RecordDelimiter *string `json:"record_delimiter,omitempty"`

	// **参数解释**： 存储Key。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	StoreKeys *bool `json:"store_keys,omitempty"`

	// **参数解释**： 每个传输文件多大后就开始上传，单位为byte；默认值5242880。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	ObsPartSize *int32 `json:"obs_part_size,omitempty"`

	// **参数解释**： 刷写数量。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	FlushSize *int32 `json:"flush_size,omitempty"`

	// **参数解释**： 时区。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	Timezone *string `json:"timezone,omitempty"`

	// **参数解释**： schema_generator类，默认为\"io.confluent.connect.storage.hive.schema.DefaultSchemaGenerator\"。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	SchemaGeneratorClass *string `json:"schema_generator_class,omitempty"`

	// **参数解释**： partitioner类，默认\"io.confluent.connect.storage.partitioner.TimeBasedPartitioner\"。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	PartitionerClass *string `json:"partitioner_class,omitempty"`

	// **参数解释**： 值转换器，默认为\"org.apache.kafka.connect.converters.ByteArrayConverter\"。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	ValueConverter *string `json:"value_converter,omitempty"`

	// **参数解释**： 键转换器，默认为\"org.apache.kafka.connect.converters.ByteArrayConverter\"。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	KeyConverter *string `json:"key_converter,omitempty"`

	// **参数解释**： 键值分隔符，默认为\":\"。（仅目标端类型为OBS时会显示） **取值范围**： 不涉及。
	KvDelimiter *string `json:"kv_delimiter,omitempty"`
}

func (o SmartConnectTaskRespSinkConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartConnectTaskRespSinkConfig struct{}"
	}

	return strings.Join([]string{"SmartConnectTaskRespSinkConfig", string(data)}, " ")
}
