package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SystemInfo struct {

	// CPU使用量
	CpuUse *float64 `json:"cpu_use,omitempty"`

	// 记录ID
	Id *string `json:"id,omitempty"`

	// 内存使用量
	MemUse *float64 `json:"mem_use,omitempty"`

	// 记录时间
	Time *string `json:"time,omitempty"`
}

func (o SystemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemInfo struct{}"
	}

	return strings.Join([]string{"SystemInfo", string(data)}, " ")
}
