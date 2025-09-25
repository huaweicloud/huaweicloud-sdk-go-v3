package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTimeDetails 资源耗时详细信息
type ResourceTimeDetails struct {

	// **参数解释**: CPU时间（单位：微秒）。 **取值范围**: 不涉及。
	CpuTime int64 `json:"cpu_time"`

	// **参数解释**: IO上的时间花费（单位：微秒）。 **取值范围**: 不涉及。
	DataIoTime int64 `json:"data_io_time"`

	// **参数解释**: 其余耗时（单位：微秒）。 **取值范围**: 不涉及。
	OtherTime int64 `json:"other_time"`
}

func (o ResourceTimeDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTimeDetails struct{}"
	}

	return strings.Join([]string{"ResourceTimeDetails", string(data)}, " ")
}
