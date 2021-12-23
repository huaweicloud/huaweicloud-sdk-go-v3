package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 复制对元数据
type ReplicationRecordMetadata struct {
	// 复制对中的云硬盘是否为共享云硬盘。

	Multiattach bool `json:"multiattach"`
	// 复制对中的云硬盘是否为系统盘。

	Bootable bool `json:"bootable"`
	// 复制对中的云硬盘容量。单位：GB

	VolumeSize int32 `json:"volume_size"`
	// 复制对中的云硬盘类型。SATA：普通IO磁盘类型。SAS：高IO磁盘类型。SSD：超高IO磁盘类型。co-p1：高IO（性能优化I型）uh-l1：超高IO（时延优化）其中co-p1和uh-l1两种云硬盘只能使用在HANA云服务器、HL1型云服务器、HL2型云服务器上。

	VolumeType ReplicationRecordMetadataVolumeType `json:"volume_type"`
}

func (o ReplicationRecordMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationRecordMetadata struct{}"
	}

	return strings.Join([]string{"ReplicationRecordMetadata", string(data)}, " ")
}

type ReplicationRecordMetadataVolumeType struct {
	value string
}

type ReplicationRecordMetadataVolumeTypeEnum struct {
	SATA  ReplicationRecordMetadataVolumeType
	SAS   ReplicationRecordMetadataVolumeType
	SSD   ReplicationRecordMetadataVolumeType
	CO_P1 ReplicationRecordMetadataVolumeType
	UH_L1 ReplicationRecordMetadataVolumeType
}

func GetReplicationRecordMetadataVolumeTypeEnum() ReplicationRecordMetadataVolumeTypeEnum {
	return ReplicationRecordMetadataVolumeTypeEnum{
		SATA: ReplicationRecordMetadataVolumeType{
			value: "SATA",
		},
		SAS: ReplicationRecordMetadataVolumeType{
			value: "SAS",
		},
		SSD: ReplicationRecordMetadataVolumeType{
			value: "SSD",
		},
		CO_P1: ReplicationRecordMetadataVolumeType{
			value: "co-p1",
		},
		UH_L1: ReplicationRecordMetadataVolumeType{
			value: "uh-l1",
		},
	}
}

func (c ReplicationRecordMetadataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReplicationRecordMetadataVolumeType) UnmarshalJSON(b []byte) error {
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
