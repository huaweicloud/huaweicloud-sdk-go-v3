package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEcsVolume 挂在在ECS的硬盘信息
type HwcEcsVolume struct {

	// 磁盘ID，格式为UUID。
	Id string `json:"id"`

	// 删除云服务器时是否一并删除该磁盘。 true：是 false：否
	DeleteOnTermination *string `json:"delete_on_termination,omitempty"`

	// 云硬盘启动顺序。 0为系统盘。 非0为数据盘。
	BootIndex *string `json:"boot_index,omitempty"`

	// 云硬盘挂载盘符，即磁盘挂载点。
	Device *string `json:"device,omitempty"`
}

func (o HwcEcsVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEcsVolume struct{}"
	}

	return strings.Join([]string{"HwcEcsVolume", string(data)}, " ")
}
