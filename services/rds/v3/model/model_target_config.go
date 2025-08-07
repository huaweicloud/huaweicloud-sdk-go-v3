package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TargetConfig struct {

	// 当name=RESIZE_FLAVOR时，表示规格变更任务的目标规格。
	Flavor *string `json:"flavor,omitempty"`

	// 当name=RESIZE_FLAVOR时，表示规格变更任务的的目标cpu。
	Cpu *string `json:"cpu,omitempty"`

	// 当name=RESIZE_FLAVOR时，表示规格变更任务的目标内存。
	Mem *string `json:"mem,omitempty"`
}

func (o TargetConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetConfig struct{}"
	}

	return strings.Join([]string{"TargetConfig", string(data)}, " ")
}
