package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeAttach 云服务器挂载磁盘信息。
type VolumeAttach struct {

	// 云磁盘ID。
	Id string `json:"id"`

	// 一个标志，指示在删除服务器时是否删除附加的卷。 默认情况下，这是False
	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`

	// 挂载点
	Device *string `json:"device,omitempty"`

	// 盘在云服务器上的挂载顺序，0表示启动盘。
	BootIndex *string `json:"bootIndex,omitempty"`

	// 云盘大小（单位：GB）。
	Size *int32 `json:"size,omitempty"`
}

func (o VolumeAttach) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeAttach struct{}"
	}

	return strings.Join([]string{"VolumeAttach", string(data)}, " ")
}
