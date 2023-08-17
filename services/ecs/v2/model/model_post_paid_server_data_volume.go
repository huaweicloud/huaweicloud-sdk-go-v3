package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PostPaidServerDataVolume 云服务器对应数据盘相关配置。
type PostPaidServerDataVolume struct {

	// 云服务器数据盘对应的磁盘类型，需要与系统所提供的磁盘类型相匹配。   - SATA：普通IO磁盘类型。  - SAS：高IO磁盘类型。  - SSD：超高IO磁盘类型。  - co-p1：高IO (性能优化Ⅰ型)。  - uh-l1：超高IO (时延优化)。  - GPSSD2：通用型SSD V2磁盘类型。  - ESSD2：极速型SSD V2磁盘类型。   > 说明： >  > 对于HANA云服务器、HL1型云服务器、HL2型云服务器，需使用co-p1和uh-l1两种磁盘类型。对于其他类型的云服务器，不能使用co-p1和uh-l1两种磁盘类型。  了解不同磁盘类型的详细信息，请参见 [磁盘类型及性能介绍](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)。  获取region可用的卷类型，请参见[查询云硬盘类型列表](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=EVS&api=CinderListVolumeTypes)
	Volumetype PostPaidServerDataVolumeVolumetype `json:"volumetype"`

	// 数据盘大小，容量单位为GB，输入大小范围为[10,32768]。
	Size int32 `json:"size"`

	// 给云硬盘配置iops，购买GPSSD2、ESSD2类型的云硬盘时必填，其他类型不能设置。  说明： 1、了解GPSSD2、ESSD2类型的iops大小范围，请参见 [云硬盘类型及性能介绍里面的云硬盘性能数据表](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)。 2、只支持按需计费。
	Iops *int32 `json:"iops,omitempty"`

	// 给云硬盘配置吞吐量，单位是MiB/s，购买GPSSD2类型云盘时必填，其他类型不能设置。  说明： 1、了解GPSSD2类型的吞吐量大小范围，请参见 [云硬盘类型及性能介绍里面的云硬盘性能数据表](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)。 2、只支持按需计费。
	Throughput *int32 `json:"throughput,omitempty"`

	// 是否为共享磁盘。true为共享盘，false为普通云硬盘。  > 说明： >  > 该字段已废弃，请使用multiattach。
	Shareable *bool `json:"shareable,omitempty"`

	// 创建共享磁盘的信息。  - true：创建的磁盘为共享盘。 - false：创建的磁盘为普通云硬盘。  > 说明： >  > shareable当前为废弃字段，如果确实需要同时使用shareable字段和multiattach字段，此时，请确保两个字段的参数值相同。当不指定该字段时，系统默认创建普通云硬盘。
	Multiattach *bool `json:"multiattach,omitempty"`

	// 数据卷是否使用SCSI锁。  - true表示云硬盘的设备类型为SCSI类型，即允许ECS操作系统直接访问底层存储介质。支持SCSI锁命令。 - false表示云硬盘的设备类型为VBD (虚拟块存储设备 , Virtual Block Device)类型，即为默认类型，VBD只能支持简单的SCSI读写命令。 - 该字段不存在时，云硬盘默认为VBD类型。  > 说明： >  > 此参数为boolean类型，若传入非boolean类型字符，程序将按照【false】方式处理。
	Hwpassthrough *bool `json:"hw:passthrough,omitempty"`

	Extendparam *PostPaidServerDataVolumeExtendParam `json:"extendparam,omitempty"`

	// 云服务器数据盘对应的磁盘存储类型。 磁盘存储类型枚举值： DSS：专属存储类型
	ClusterType *PostPaidServerDataVolumeClusterType `json:"cluster_type,omitempty"`

	// 云服务器数据盘对应的存储池的ID。
	ClusterId *string `json:"cluster_id,omitempty"`

	Metadata *PostPaidServerDataVolumeMetadata `json:"metadata,omitempty"`

	// 数据镜像的ID，UUID格式。  如果使用数据盘镜像创建数据盘，则data_image_id为必选参数，且不支持使用metadata。
	DataImageId *string `json:"data_image_id,omitempty"`

	// 数据盘随实例释放策略。  true：数据盘随实例释放。 false：数据盘不随实例释放。 默认值：false。
	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`
}

func (o PostPaidServerDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerDataVolume struct{}"
	}

	return strings.Join([]string{"PostPaidServerDataVolume", string(data)}, " ")
}

type PostPaidServerDataVolumeVolumetype struct {
	value string
}

type PostPaidServerDataVolumeVolumetypeEnum struct {
	SATA   PostPaidServerDataVolumeVolumetype
	SAS    PostPaidServerDataVolumeVolumetype
	SSD    PostPaidServerDataVolumeVolumetype
	GPSSD  PostPaidServerDataVolumeVolumetype
	CO_P1  PostPaidServerDataVolumeVolumetype
	UH_L1  PostPaidServerDataVolumeVolumetype
	ESSD   PostPaidServerDataVolumeVolumetype
	GPSSD2 PostPaidServerDataVolumeVolumetype
	ESSD2  PostPaidServerDataVolumeVolumetype
}

func GetPostPaidServerDataVolumeVolumetypeEnum() PostPaidServerDataVolumeVolumetypeEnum {
	return PostPaidServerDataVolumeVolumetypeEnum{
		SATA: PostPaidServerDataVolumeVolumetype{
			value: "SATA",
		},
		SAS: PostPaidServerDataVolumeVolumetype{
			value: "SAS",
		},
		SSD: PostPaidServerDataVolumeVolumetype{
			value: "SSD",
		},
		GPSSD: PostPaidServerDataVolumeVolumetype{
			value: "GPSSD",
		},
		CO_P1: PostPaidServerDataVolumeVolumetype{
			value: "co-p1",
		},
		UH_L1: PostPaidServerDataVolumeVolumetype{
			value: "uh-l1",
		},
		ESSD: PostPaidServerDataVolumeVolumetype{
			value: "ESSD",
		},
		GPSSD2: PostPaidServerDataVolumeVolumetype{
			value: "GPSSD2",
		},
		ESSD2: PostPaidServerDataVolumeVolumetype{
			value: "ESSD2",
		},
	}
}

func (c PostPaidServerDataVolumeVolumetype) Value() string {
	return c.value
}

func (c PostPaidServerDataVolumeVolumetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerDataVolumeVolumetype) UnmarshalJSON(b []byte) error {
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

type PostPaidServerDataVolumeClusterType struct {
	value string
}

type PostPaidServerDataVolumeClusterTypeEnum struct {
	DSS PostPaidServerDataVolumeClusterType
}

func GetPostPaidServerDataVolumeClusterTypeEnum() PostPaidServerDataVolumeClusterTypeEnum {
	return PostPaidServerDataVolumeClusterTypeEnum{
		DSS: PostPaidServerDataVolumeClusterType{
			value: "DSS",
		},
	}
}

func (c PostPaidServerDataVolumeClusterType) Value() string {
	return c.value
}

func (c PostPaidServerDataVolumeClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerDataVolumeClusterType) UnmarshalJSON(b []byte) error {
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
