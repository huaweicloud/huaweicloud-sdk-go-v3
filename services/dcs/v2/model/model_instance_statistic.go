package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例的统计信息。
type InstanceStatistic struct {

	// 缓存实例网络入流量，单位：Kbps。
	InputKbps *string `json:"input_kbps,omitempty" xml:"input_kbps"`

	// 缓存实例网络出流量，单位：Kbps。
	OutputKbps *string `json:"output_kbps,omitempty" xml:"output_kbps"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 缓存存储的数据条数。
	Keys *int64 `json:"keys,omitempty" xml:"keys"`

	// 缓存已经使用内存，单位：MB。
	UsedMemory *int64 `json:"used_memory,omitempty" xml:"used_memory"`

	// 缓存的总内存，单位：MB。
	MaxMemory *int64 `json:"max_memory,omitempty" xml:"max_memory"`

	// 缓存get命令被调用次数。
	CmdGetCount *int64 `json:"cmd_get_count,omitempty" xml:"cmd_get_count"`

	// 缓存set命令被调用次数。
	CmdSetCount *int64 `json:"cmd_set_count,omitempty" xml:"cmd_set_count"`

	// CPU使用率，单位：百分比。
	UsedCpu *string `json:"used_cpu,omitempty" xml:"used_cpu"`
}

func (o InstanceStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceStatistic struct{}"
	}

	return strings.Join([]string{"InstanceStatistic", string(data)}, " ")
}
