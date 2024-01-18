package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StorageMetadata 存储定义。
type StorageMetadata struct {

	// SFS文件系统名称。
	StorageHandle string `json:"storage_handle"`

	// 存储类型： * `sfs` - sfs3.0存储。
	StorageClass string `json:"storage_class"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// 所在区域。
	Region *string `json:"region,omitempty"`

	// 访问地址:protocol://[bucket-name].sfs3.[region-name].myhuaweicloud.com:port。
	ExportLocation *string `json:"export_location,omitempty"`
}

func (o StorageMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageMetadata struct{}"
	}

	return strings.Join([]string{"StorageMetadata", string(data)}, " ")
}
