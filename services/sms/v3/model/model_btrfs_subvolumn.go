package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BtrfsSubvolumn btrfs子卷信息
type BtrfsSubvolumn struct {

	// 父卷的uuid
	Uuid *string `json:"uuid,omitempty"`

	// 子卷是否为快照
	IsSnapshot *string `json:"is_snapshot,omitempty"`

	// 子卷的ID
	SubvolId *string `json:"subvol_id,omitempty"`

	// 父卷ID
	ParentId *string `json:"parent_id,omitempty"`

	// 子卷的名称
	SubvolName *string `json:"subvol_name,omitempty"`

	// 子卷的挂载路径
	SubvolMountPath *string `json:"subvol_mount_path,omitempty"`
}

func (o BtrfsSubvolumn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BtrfsSubvolumn struct{}"
	}

	return strings.Join([]string{"BtrfsSubvolumn", string(data)}, " ")
}
