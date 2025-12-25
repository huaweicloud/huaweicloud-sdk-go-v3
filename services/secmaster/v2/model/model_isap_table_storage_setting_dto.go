package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapTableStorageSettingDto 表存储设置数据传输对象
type IsapTableStorageSettingDto struct {

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

	// 流式数据空间ID
	StreamingDataspaceId *string `json:"streaming_dataspace_id,omitempty"`

	// 索引存储周期
	IndexStoragePeriod *int32 `json:"index_storage_period,omitempty"`

	// 索引存储大小
	IndexStorageSize *int32 `json:"index_storage_size,omitempty"`

	// 索引分片数
	IndexShards *int32 `json:"index_shards,omitempty"`

	// 数据湖存储周期
	LakeStoragePeriod *int64 `json:"lake_storage_period,omitempty"`

	// **参数解释**: 数据湖定时分区设置 - MINUTE10 10分钟 - HOUR 小时 - DAY 天  **约束限制** 不涉及 **取值范围**: - MINUTE - HOUR - DAY  **默认值** 不涉及
	LakePartitionSetting *IsapTableStorageSettingDtoLakePartitionSetting `json:"lake_partition_setting,omitempty"`
}

func (o IsapTableStorageSettingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTableStorageSettingDto struct{}"
	}

	return strings.Join([]string{"IsapTableStorageSettingDto", string(data)}, " ")
}

type IsapTableStorageSettingDtoLakePartitionSetting struct {
	value string
}

type IsapTableStorageSettingDtoLakePartitionSettingEnum struct {
	DAY      IsapTableStorageSettingDtoLakePartitionSetting
	HOUR     IsapTableStorageSettingDtoLakePartitionSetting
	MINUTE10 IsapTableStorageSettingDtoLakePartitionSetting
}

func GetIsapTableStorageSettingDtoLakePartitionSettingEnum() IsapTableStorageSettingDtoLakePartitionSettingEnum {
	return IsapTableStorageSettingDtoLakePartitionSettingEnum{
		DAY: IsapTableStorageSettingDtoLakePartitionSetting{
			value: "DAY",
		},
		HOUR: IsapTableStorageSettingDtoLakePartitionSetting{
			value: "HOUR",
		},
		MINUTE10: IsapTableStorageSettingDtoLakePartitionSetting{
			value: "MINUTE10",
		},
	}
}

func (c IsapTableStorageSettingDtoLakePartitionSetting) Value() string {
	return c.value
}

func (c IsapTableStorageSettingDtoLakePartitionSetting) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapTableStorageSettingDtoLakePartitionSetting) UnmarshalJSON(b []byte) error {
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
