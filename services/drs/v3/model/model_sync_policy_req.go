package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 设置同步策略请求体
type SyncPolicyReq struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 冲突策略。 - ignore：忽略 - overwrite：覆盖 - stop：报错
	ConflictPolicy *SyncPolicyReqConflictPolicy `json:"conflict_policy,omitempty"`

	// 过滤DDL策略。
	FilterDdlPolicy *SyncPolicyReqFilterDdlPolicy `json:"filter_ddl_policy,omitempty"`

	// 同步增量是否同步DDL。
	DdlTrans *bool `json:"ddl_trans,omitempty"`

	// 同步增量是否同步索引。
	IndexTrans *bool `json:"index_trans,omitempty"`

	// 同步Topic策略，目标库为kafka时必填，取值： - 0：集中投递到一个Topic - 1：按库名-schema-表名自动生成Topic名字 - 2：按库名自动生成Topic名字 - 3：按库名-schema自动生成Topic名字
	TopicPolicy *SyncPolicyReqTopicPolicy `json:"topic_policy,omitempty"`

	// Topic名称，topic_policy为0时必填，确保topic已存在。
	Topic *string `json:"topic,omitempty"`

	// 同步到kafka partition策略，取值： - 0：按库名.schema.表名的hash值投递到不同Partition - 1：全部投递到Partition 0 - 2：按主键的hash值投递到不同Partition - 3：按库名.schema的hash值投递到不同Partition **当topic_policy取0时，可以取0,1,2,3；当topic_policy取1时，可以取1,2；当topic_policy取2时，可以取0,1,3；当topic_policy取3时，可以取0,1；**
	PartitionPolicy *SyncPolicyReqPartitionPolicy `json:"partition_policy,omitempty"`

	// 投送到kafka的数据格式，不填默认为json：
	KafkaDataFormat *SyncPolicyReqKafkaDataFormat `json:"kafka_data_format,omitempty"`

	// Topic名字格式，topic_policy为1,2,3,时需要 - 当topic_policy取1时，Topic名字格式支持database、schema两个变量，其他字符当做常量。分别用$database$代替数据库名，$schema$代替模式名，不填默认为$database$-$schema$ - 当topic_policy取2时，Topic名字格式支持database一个变量，其他字符都当做常量，不填默认为$database$ - 当topic_policy取3时，Topic名字格式支持database、schema和tablename三个变量，其他字符当做常量。分别用$database$代替数据库名，$schema$代替模式名，$tablename$代替表名，不填默认为$database$-$schema$-$tablename$
	TopicNameFormat *string `json:"topic_name_format,omitempty"`

	// Partition个数，取值1-2147483647之间，topic_policy为1,2,3,时需要，不填默认为1
	PartitionsNum *string `json:"partitions_num,omitempty"`

	// 副本个数，取值1-32767之间，topic_policy为1,2,3,时需要，不填默认为1
	ReplicationFactor *string `json:"replication_factor,omitempty"`

	// PostgreSQL同步全量阶段是否填充物化视图，不填默认为false
	IsFillMaterializedView *bool `json:"is_fill_materialized_view,omitempty"`

	// PostgreSQL同步全量阶段是否使用快照模式导出，不填默认为false
	ExportSnapshot *bool `json:"export_snapshot,omitempty"`

	// 复制槽名称，gaussdbv5ha-to-kafka主备任务必填
	SlotName *string `json:"slot_name,omitempty"`
}

func (o SyncPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncPolicyReq struct{}"
	}

	return strings.Join([]string{"SyncPolicyReq", string(data)}, " ")
}

type SyncPolicyReqConflictPolicy struct {
	value string
}

type SyncPolicyReqConflictPolicyEnum struct {
	IGNORE    SyncPolicyReqConflictPolicy
	OVERWRITE SyncPolicyReqConflictPolicy
	STOP      SyncPolicyReqConflictPolicy
}

func GetSyncPolicyReqConflictPolicyEnum() SyncPolicyReqConflictPolicyEnum {
	return SyncPolicyReqConflictPolicyEnum{
		IGNORE: SyncPolicyReqConflictPolicy{
			value: "ignore",
		},
		OVERWRITE: SyncPolicyReqConflictPolicy{
			value: "overwrite",
		},
		STOP: SyncPolicyReqConflictPolicy{
			value: "stop",
		},
	}
}

func (c SyncPolicyReqConflictPolicy) Value() string {
	return c.value
}

func (c SyncPolicyReqConflictPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqConflictPolicy) UnmarshalJSON(b []byte) error {
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

type SyncPolicyReqFilterDdlPolicy struct {
	value string
}

type SyncPolicyReqFilterDdlPolicyEnum struct {
	DROP_DATABASE SyncPolicyReqFilterDdlPolicy
}

func GetSyncPolicyReqFilterDdlPolicyEnum() SyncPolicyReqFilterDdlPolicyEnum {
	return SyncPolicyReqFilterDdlPolicyEnum{
		DROP_DATABASE: SyncPolicyReqFilterDdlPolicy{
			value: "drop_database",
		},
	}
}

func (c SyncPolicyReqFilterDdlPolicy) Value() string {
	return c.value
}

func (c SyncPolicyReqFilterDdlPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqFilterDdlPolicy) UnmarshalJSON(b []byte) error {
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

type SyncPolicyReqTopicPolicy struct {
	value string
}

type SyncPolicyReqTopicPolicyEnum struct {
	E_0 SyncPolicyReqTopicPolicy
	E_1 SyncPolicyReqTopicPolicy
	E_2 SyncPolicyReqTopicPolicy
	E_3 SyncPolicyReqTopicPolicy
}

func GetSyncPolicyReqTopicPolicyEnum() SyncPolicyReqTopicPolicyEnum {
	return SyncPolicyReqTopicPolicyEnum{
		E_0: SyncPolicyReqTopicPolicy{
			value: "0",
		},
		E_1: SyncPolicyReqTopicPolicy{
			value: "1",
		},
		E_2: SyncPolicyReqTopicPolicy{
			value: "2",
		},
		E_3: SyncPolicyReqTopicPolicy{
			value: "3",
		},
	}
}

func (c SyncPolicyReqTopicPolicy) Value() string {
	return c.value
}

func (c SyncPolicyReqTopicPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqTopicPolicy) UnmarshalJSON(b []byte) error {
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

type SyncPolicyReqPartitionPolicy struct {
	value string
}

type SyncPolicyReqPartitionPolicyEnum struct {
	E_0 SyncPolicyReqPartitionPolicy
	E_1 SyncPolicyReqPartitionPolicy
	E_2 SyncPolicyReqPartitionPolicy
	E_3 SyncPolicyReqPartitionPolicy
}

func GetSyncPolicyReqPartitionPolicyEnum() SyncPolicyReqPartitionPolicyEnum {
	return SyncPolicyReqPartitionPolicyEnum{
		E_0: SyncPolicyReqPartitionPolicy{
			value: "0",
		},
		E_1: SyncPolicyReqPartitionPolicy{
			value: "1",
		},
		E_2: SyncPolicyReqPartitionPolicy{
			value: "2",
		},
		E_3: SyncPolicyReqPartitionPolicy{
			value: "3",
		},
	}
}

func (c SyncPolicyReqPartitionPolicy) Value() string {
	return c.value
}

func (c SyncPolicyReqPartitionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqPartitionPolicy) UnmarshalJSON(b []byte) error {
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

type SyncPolicyReqKafkaDataFormat struct {
	value string
}

type SyncPolicyReqKafkaDataFormatEnum struct {
	JSON SyncPolicyReqKafkaDataFormat
	AVRO SyncPolicyReqKafkaDataFormat
}

func GetSyncPolicyReqKafkaDataFormatEnum() SyncPolicyReqKafkaDataFormatEnum {
	return SyncPolicyReqKafkaDataFormatEnum{
		JSON: SyncPolicyReqKafkaDataFormat{
			value: "json",
		},
		AVRO: SyncPolicyReqKafkaDataFormat{
			value: "avro",
		},
	}
}

func (c SyncPolicyReqKafkaDataFormat) Value() string {
	return c.value
}

func (c SyncPolicyReqKafkaDataFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqKafkaDataFormat) UnmarshalJSON(b []byte) error {
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
