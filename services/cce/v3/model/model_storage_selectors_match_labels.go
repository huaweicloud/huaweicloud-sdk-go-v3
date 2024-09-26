package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StorageSelectorsMatchLabels evs盘的匹配字段，支持DataVolume中的size、volumeType、[iops、throughput、](tag:hws)metadataEncrypted、metadataCmkid、count字段。
type StorageSelectorsMatchLabels struct {

	// 匹配的磁盘大小，不填则无磁盘大小限制。例如：100.
	Size *string `json:"size,omitempty"`

	// 云硬盘类型，目前支持SSD\\GPSSD\\SAS\\ESSD\\SATA等。
	VolumeType *string `json:"volumeType,omitempty"`

	// 匹配的磁盘iops大小，不填则无磁盘iops大小限制。当需要选择GPSSD2或ESSD2类型磁盘时，配置iops来准确选择磁盘。例如：3000.
	Iops *string `json:"iops,omitempty"`

	// 匹配的磁盘吞吐量大小，不填则无磁盘吞吐量大小限制。当需要选择GPSSD2类型磁盘时，配置throughput来准确选择磁盘。例如：125.
	Throughput *string `json:"throughput,omitempty"`

	// 磁盘加密标识符，0代表不加密，1代表加密。
	MetadataEncrypted *string `json:"metadataEncrypted,omitempty"`

	// 加密磁盘的用户主密钥ID，长度为36字节的字符串。
	MetadataCmkid *string `json:"metadataCmkid,omitempty"`

	// 磁盘选择个数，不填则选择所有此类磁盘。
	Count *string `json:"count,omitempty"`
}

func (o StorageSelectorsMatchLabels) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageSelectorsMatchLabels struct{}"
	}

	return strings.Join([]string{"StorageSelectorsMatchLabels", string(data)}, " ")
}
