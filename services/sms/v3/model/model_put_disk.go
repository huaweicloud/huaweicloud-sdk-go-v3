package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 磁盘信息
type PutDisk struct {

	// 磁盘名称
	NeedMigration *bool `json:"need_migration,omitempty" xml:"need_migration"`

	// 磁盘id
	Id string `json:"id" xml:"id"`

	// 调整大小
	AdjustSize int64 `json:"adjust_size" xml:"adjust_size"`

	// 修改的卷信息
	PhysicalVolumes *[]PutVolume `json:"physical_volumes,omitempty" xml:"physical_volumes"`
}

func (o PutDisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutDisk struct{}"
	}

	return strings.Join([]string{"PutDisk", string(data)}, " ")
}
