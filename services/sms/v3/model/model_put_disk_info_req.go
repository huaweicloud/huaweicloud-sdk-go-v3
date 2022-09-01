package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type PutDiskInfoReq struct {

	// 更新的磁盘信息
	Disks *[]ServerDisk `json:"disks,omitempty" xml:"disks"`

	// 更新的卷信息
	Volumegroups *[]VolumeGroups `json:"volumegroups,omitempty" xml:"volumegroups"`

	// 更新的btrfs信息
	BtrfsList *[]BtrfsFileSystem `json:"btrfs_list,omitempty" xml:"btrfs_list"`
}

func (o PutDiskInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutDiskInfoReq struct{}"
	}

	return strings.Join([]string{"PutDiskInfoReq", string(data)}, " ")
}
