package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PostPaidServerRootVolume
type PostPaidServerRootVolume struct {

	// 云服务器系统盘对应的磁盘类型，需要与系统所提供的磁盘类型相匹配。   - SATA：普通IO磁盘类型。  - SAS：高IO磁盘类型。  - SSD：超高IO磁盘类型。  - GPSSD：为通用型SSD磁盘类型。  - co-p1：高IO (性能优化Ⅰ型)。  - uh-l1：超高IO (时延优化)。  - ESSD：为极速IO磁盘类型。  - GPSSD2：通用型SSD V2磁盘类型。  - ESSD2：极速型SSD V2磁盘类型。   > 说明： >  > 对于HANA云服务器、HL1型云服务器、HL2型云服务器，需使用co-p1和uh-l1两种磁盘类型。对于其他类型的云服务器，不能使用co-p1和uh-l1两种磁盘类型。  了解不同磁盘类型的详细信息，请参见 [磁盘类型及性能介绍](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)。  获取region可用的卷类型，请参见[查询云硬盘类型列表](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=EVS&api=CinderListVolumeTypes)
	Volumetype PostPaidServerRootVolumeVolumetype `json:"volumetype"`

	// 系统盘大小，容量单位为GB， 输入大小范围为[1,1024]。  约束：  - 系统盘大小取值应不小于镜像支持的系统盘的最小值(镜像的min_disk属性)。 - 若该参数没有指定或者指定为0时，系统盘大小默认取值为镜像中系统盘的最小值(镜像的min_disk属性)。  > 说明： > > 镜像系统盘的最小值(镜像的min_disk属性)可在控制台中点击镜像详情查看。或通过调用“查询镜像详情（OpenStack原生）”API获取，详细操作请参考[《镜像服务API参考》](https://support.huaweicloud.com/api-ims/ims_03_0702.html)中“查询镜像详情（OpenStack原生）”章节。
	Size *int32 `json:"size,omitempty"`

	// 给云硬盘配置iops，购买GPSSD2、ESSD2类型的云硬盘时必填，其他类型不能设置。  说明： 1、了解GPSSD2、ESSD2类型的iops大小范围，请参见 [云硬盘类型及性能介绍里面的云硬盘性能数据表](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)。 2、只支持按需计费。
	Iops *int32 `json:"iops,omitempty"`

	// 给云硬盘配置吞吐量，单位是MiB/s，购买GPSSD2类型云盘时必填，其他类型不能设置。  说明： 1、了解GPSSD2类型的吞吐量大小范围，请参见 [云硬盘类型及性能介绍里面的云硬盘性能数据表](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)。 2、只支持按需计费。
	Throughput *int32 `json:"throughput,omitempty"`

	// 使用SDI规格创建虚拟机时请关注该参数，如果该参数值为true，说明创建的为scsi类型的卷  > 说明： >  > 此参数为boolean类型，若传入非boolean类型字符，程序将按照false方式处理。
	Hwpassthrough *bool `json:"hw:passthrough,omitempty"`

	// 云服务器系统盘对应的磁盘存储类型。 磁盘存储类型枚举值： DSS：专属存储类型
	ClusterType *PostPaidServerRootVolumeClusterType `json:"cluster_type,omitempty"`

	// 云服务器数据盘对应的存储池的ID。
	ClusterId *string `json:"cluster_id,omitempty"`

	Extendparam *PostPaidServerRootVolumeExtendParam `json:"extendparam,omitempty"`

	Metadata *PostPaidServerRootVolumeMetadata `json:"metadata,omitempty"`
}

func (o PostPaidServerRootVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerRootVolume struct{}"
	}

	return strings.Join([]string{"PostPaidServerRootVolume", string(data)}, " ")
}

type PostPaidServerRootVolumeVolumetype struct {
	value string
}

type PostPaidServerRootVolumeVolumetypeEnum struct {
	SATA   PostPaidServerRootVolumeVolumetype
	SAS    PostPaidServerRootVolumeVolumetype
	SSD    PostPaidServerRootVolumeVolumetype
	GPSSD  PostPaidServerRootVolumeVolumetype
	CO_P1  PostPaidServerRootVolumeVolumetype
	UH_L1  PostPaidServerRootVolumeVolumetype
	ESSD   PostPaidServerRootVolumeVolumetype
	GPSSD2 PostPaidServerRootVolumeVolumetype
	ESSD2  PostPaidServerRootVolumeVolumetype
}

func GetPostPaidServerRootVolumeVolumetypeEnum() PostPaidServerRootVolumeVolumetypeEnum {
	return PostPaidServerRootVolumeVolumetypeEnum{
		SATA: PostPaidServerRootVolumeVolumetype{
			value: "SATA",
		},
		SAS: PostPaidServerRootVolumeVolumetype{
			value: "SAS",
		},
		SSD: PostPaidServerRootVolumeVolumetype{
			value: "SSD",
		},
		GPSSD: PostPaidServerRootVolumeVolumetype{
			value: "GPSSD",
		},
		CO_P1: PostPaidServerRootVolumeVolumetype{
			value: "co-p1",
		},
		UH_L1: PostPaidServerRootVolumeVolumetype{
			value: "uh-l1",
		},
		ESSD: PostPaidServerRootVolumeVolumetype{
			value: "ESSD",
		},
		GPSSD2: PostPaidServerRootVolumeVolumetype{
			value: "GPSSD2",
		},
		ESSD2: PostPaidServerRootVolumeVolumetype{
			value: "ESSD2",
		},
	}
}

func (c PostPaidServerRootVolumeVolumetype) Value() string {
	return c.value
}

func (c PostPaidServerRootVolumeVolumetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerRootVolumeVolumetype) UnmarshalJSON(b []byte) error {
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

type PostPaidServerRootVolumeClusterType struct {
	value string
}

type PostPaidServerRootVolumeClusterTypeEnum struct {
	DSS PostPaidServerRootVolumeClusterType
}

func GetPostPaidServerRootVolumeClusterTypeEnum() PostPaidServerRootVolumeClusterTypeEnum {
	return PostPaidServerRootVolumeClusterTypeEnum{
		DSS: PostPaidServerRootVolumeClusterType{
			value: "DSS",
		},
	}
}

func (c PostPaidServerRootVolumeClusterType) Value() string {
	return c.value
}

func (c PostPaidServerRootVolumeClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerRootVolumeClusterType) UnmarshalJSON(b []byte) error {
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
