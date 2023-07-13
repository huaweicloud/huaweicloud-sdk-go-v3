package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Storage 存储定义元数据
type Storage struct {

	// SFS文件系统名称
	StorageHandle string `json:"storage_handle"`

	// 存储类型 * `sfs` - sfs3.0存储
	StorageClass string `json:"storage_class"`
}

func (o Storage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Storage struct{}"
	}

	return strings.Join([]string{"Storage", string(data)}, " ")
}
