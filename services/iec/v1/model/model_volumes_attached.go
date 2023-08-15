package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumesAttached
type VolumesAttached struct {

	// 磁盘ID，格式为UUID。
	Id *string `json:"id,omitempty"`

	// 启动标识。  - “0”代表系统盘 - 非“0”为数据盘。
	BootIndex *string `json:"bootIndex,omitempty"`

	// 删边缘实例时是否一并删除该磁盘。  - true：是 - false：否
	DeleteOnTermination *string `json:"delete_on_termination,omitempty"`

	// 硬盘挂载盘符，即磁盘挂载点。
	Device *string `json:"device,omitempty"`
}

func (o VolumesAttached) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumesAttached struct{}"
	}

	return strings.Join([]string{"VolumesAttached", string(data)}, " ")
}
