package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BtrfsSubvolumn btrfs子卷信息
type BtrfsSubvolumn struct {

	// 父卷的uuid
	Uuid string `json:"uuid"`

	// 子卷是否为快照
	IsSnapshot string `json:"is_snapshot"`

	// 子卷的ID
	SubvolId string `json:"subvol_id"`

	// 父卷ID
	ParentId string `json:"parent_id"`

	// 子卷的名称
	SubvolName string `json:"subvol_name"`

	// 子卷的挂载路径
	SubvolMountPath string `json:"subvol_mount_path"`
}

func (o BtrfsSubvolumn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BtrfsSubvolumn struct{}"
	}

	return strings.Join([]string{"BtrfsSubvolumn", string(data)}, " ")
}
