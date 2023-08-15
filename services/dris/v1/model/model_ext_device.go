package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtDevice 对应linux上device概念，用于串口、AI加速卡的挂载
type ExtDevice struct {

	// **参数说明**：源路径。
	Source string `json:"source"`

	// **参数说明**：卷挂载路径。
	Destination string `json:"destination"`

	// **参数说明**：只读，默认MRW。
	CgroupPermissions *string `json:"cgroup_permissions,omitempty"`
}

func (o ExtDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtDevice struct{}"
	}

	return strings.Join([]string{"ExtDevice", string(data)}, " ")
}
