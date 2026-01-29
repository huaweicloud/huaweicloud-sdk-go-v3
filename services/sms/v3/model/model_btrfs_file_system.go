package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BtrfsFileSystem btrfs分区类型
type BtrfsFileSystem struct {

	// 文件系统名称
	Name *string `json:"name,omitempty"`

	// 文件系统标签，若无标签为空字符串
	Label *string `json:"label,omitempty"`

	// 文件系统的uuid
	Uuid *string `json:"uuid,omitempty"`

	// btrfs包含的设备名称
	Device *string `json:"device,omitempty"`

	// 文件系统数据占用大小
	Size *int64 `json:"size,omitempty"`

	// btrfs节点大小
	Nodesize *int64 `json:"nodesize,omitempty"`

	// 扇区大小
	Sectorsize *int32 `json:"sectorsize,omitempty"`

	// 数据配置（RAD）
	DataProfile *string `json:"data_profile,omitempty"`

	// 文件系统配置（RAD）
	SystemProfile *string `json:"system_profile,omitempty"`

	// 元数据配置（RAD）
	MetadataProfile *string `json:"metadata_profile,omitempty"`

	// Btrfs文件系统信息
	GlobalReserve1 *string `json:"global_reserve1,omitempty"`

	// Btrfs卷已使用空间大小
	GVolUsedSize *int64 `json:"g_vol_used_size,omitempty"`

	// 默认子卷ID
	DefaultSubvolid *string `json:"default_subvolid,omitempty"`

	// 默认子卷名称
	DefaultSubvolName *string `json:"default_subvol_name,omitempty"`

	// 默认子卷挂载路径/BTRFS文件系统的挂载路径
	DefaultSubvolMountpath *string `json:"default_subvol_mountpath,omitempty"`

	// 子卷信息
	Subvolumn *[]BtrfsSubvolumn `json:"subvolumn,omitempty"`
}

func (o BtrfsFileSystem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BtrfsFileSystem struct{}"
	}

	return strings.Join([]string{"BtrfsFileSystem", string(data)}, " ")
}
