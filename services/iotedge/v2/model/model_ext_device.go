package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 对应linux上device概念，用于串口、AI加速卡的挂载
type ExtDevice struct {

	// 源路径
	Source string `json:"source" xml:"source"`

	// 卷挂载路径
	Destination string `json:"destination" xml:"destination"`

	// 只读，默认MRW
	CgroupPermissions *string `json:"cgroup_permissions,omitempty" xml:"cgroup_permissions"`
}

func (o ExtDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtDevice struct{}"
	}

	return strings.Join([]string{"ExtDevice", string(data)}, " ")
}
