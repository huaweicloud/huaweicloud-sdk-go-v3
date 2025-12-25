package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableStorageSetting 表存储设置
type TableStorageSetting struct {

	// 应用索引
	ApplicationIndex *string `json:"application_index,omitempty"`

	// 应用主题
	ApplicationTopic *string `json:"application_topic,omitempty"`

	// 应用数据类ID
	ApplicationDataClassId *string `json:"application_data_class_id,omitempty"`

	// 流式带宽 MB/s
	StreamingBandwidth float32 `json:"streaming_bandwidth,omitempty"`

	// 流式分区
	StreamingPartition *int32 `json:"streaming_partition,omitempty"`

	// 流式容量大小
	StreamingRetentionSize *int32 `json:"streaming_retention_size,omitempty"`

	// 流式数据空间ID
	StreamingDataspaceId *string `json:"streaming_dataspace_id,omitempty"`

	// 索引存储周期
	IndexStoragePeriod *int32 `json:"index_storage_period,omitempty"`

	// 索引存储大小
	IndexStorageSize *int32 `json:"index_storage_size,omitempty"`

	// 索引分片数
	IndexShards *int32 `json:"index_shards,omitempty"`

	// 索引副本数
	IndexReplicas *int32 `json:"index_replicas,omitempty"`

	// 数据湖存储周期
	LakeStoragePeriod *int64 `json:"lake_storage_period,omitempty"`

	// **参数解释**: 时间单位 - MINUTE10 10分钟 - HOUR 小时 - DAY 天  **约束限制** 不涉及 **取值范围**: - MINUTE - HOUR - DAY  **默认值** 不涉及
	LakePartitionSetting *TableStorageSettingLakePartitionSetting `json:"lake_partition_setting,omitempty"`

	// **参数解释**: 数据湖过期状态 - NOT_EXPIRED 未过期 - EXPIRED_PROCESSING 过期处理中 - EXPIRED_SUCCESS 过期处理成功 - EXPIRED_FAILURE 过期处理失败  **约束限制** 不涉及 **取值范围**: - NOT_EXPIRED - EXPIRED_PROCESSING - EXPIRED_SUCCESS - EXPIRED_FAILURE  **默认值** 不涉及
	LakeExpirationStatus *TableStorageSettingLakeExpirationStatus `json:"lake_expiration_status,omitempty"`
}

func (o TableStorageSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableStorageSetting struct{}"
	}

	return strings.Join([]string{"TableStorageSetting", string(data)}, " ")
}

type TableStorageSettingLakePartitionSetting struct {
	value string
}

type TableStorageSettingLakePartitionSettingEnum struct {
	DAY      TableStorageSettingLakePartitionSetting
	HOUR     TableStorageSettingLakePartitionSetting
	MINUTE10 TableStorageSettingLakePartitionSetting
}

func GetTableStorageSettingLakePartitionSettingEnum() TableStorageSettingLakePartitionSettingEnum {
	return TableStorageSettingLakePartitionSettingEnum{
		DAY: TableStorageSettingLakePartitionSetting{
			value: "DAY",
		},
		HOUR: TableStorageSettingLakePartitionSetting{
			value: "HOUR",
		},
		MINUTE10: TableStorageSettingLakePartitionSetting{
			value: "MINUTE10",
		},
	}
}

func (c TableStorageSettingLakePartitionSetting) Value() string {
	return c.value
}

func (c TableStorageSettingLakePartitionSetting) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableStorageSettingLakePartitionSetting) UnmarshalJSON(b []byte) error {
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

type TableStorageSettingLakeExpirationStatus struct {
	value string
}

type TableStorageSettingLakeExpirationStatusEnum struct {
	NOT_EXPIRED        TableStorageSettingLakeExpirationStatus
	EXPIRED_PROCESSING TableStorageSettingLakeExpirationStatus
	EXPIRED_SUCCESS    TableStorageSettingLakeExpirationStatus
	EXPIRED_FAILURE    TableStorageSettingLakeExpirationStatus
}

func GetTableStorageSettingLakeExpirationStatusEnum() TableStorageSettingLakeExpirationStatusEnum {
	return TableStorageSettingLakeExpirationStatusEnum{
		NOT_EXPIRED: TableStorageSettingLakeExpirationStatus{
			value: "NOT_EXPIRED",
		},
		EXPIRED_PROCESSING: TableStorageSettingLakeExpirationStatus{
			value: "EXPIRED_PROCESSING",
		},
		EXPIRED_SUCCESS: TableStorageSettingLakeExpirationStatus{
			value: "EXPIRED_SUCCESS",
		},
		EXPIRED_FAILURE: TableStorageSettingLakeExpirationStatus{
			value: "EXPIRED_FAILURE",
		},
	}
}

func (c TableStorageSettingLakeExpirationStatus) Value() string {
	return c.value
}

func (c TableStorageSettingLakeExpirationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableStorageSettingLakeExpirationStatus) UnmarshalJSON(b []byte) error {
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
