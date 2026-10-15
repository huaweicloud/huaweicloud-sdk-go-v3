package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonitorDataResponse Response Object
type ShowMonitorDataResponse struct {

	// **参数解释**：  EIP带宽，用于评估公网传输速率是否满足迁移需求。单位：MB/s。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Bandwidth *string `json:"bandwidth,omitempty"`

	// **参数解释**：  源库连接状态是否正常。  **约束限制**：  不涉及。  **取值范围**：  - true：源库连接正常 - false：源库连接异常  **默认取值**：  不涉及
	IsSrcNormal *bool `json:"is_src_normal,omitempty"`

	// **参数解释**：  目标库连接状态是否正常。  **约束限制**：  不涉及。  **取值范围**：  - true：目标库连接正常 - false：目标库连接异常  **默认取值**：  不涉及
	IsDstNormal *bool `json:"is_dst_normal,omitempty"`

	// **参数解释**：  源库offset位点，用于判断源库数据同步进度。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SrcOffset *string `json:"src_offset,omitempty"`

	// **参数解释**：  迁移实例offset位点，用于判断迁移实例数据同步进度。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	NodeOffset *string `json:"node_offset,omitempty"`

	// **参数解释**：  目标库offset位点，用于判断目标库数据同步进度。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DstOffset *string `json:"dst_offset,omitempty"`

	// **参数解释**：  源库时延，用于判断数据同步的实时性。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SrcDelay *int32 `json:"src_delay,omitempty"`

	// **参数解释**：  目标库时延，用于判断数据同步的实时性。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DstDelay *int32 `json:"dst_delay,omitempty"`

	// **参数解释**：  源库RPS（Requests Per Second，每秒请求数），用于评估源库的数据变更速率。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SrcRps *string `json:"src_rps,omitempty"`

	// **参数解释**：  源库IO（Input/Output，每秒输入输出操作量），用于评估源库的磁盘读写负载。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SrcIo *string `json:"src_io,omitempty"`

	// **参数解释**：  目标库RPS（Requests Per Second，每秒请求数），用于评估目标库的数据写入速率。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DstRps *string `json:"dst_rps,omitempty"`

	// **参数解释**：  目标库IO（Input/Output，每秒输入输出操作量），用于评估目标库的磁盘读写负载。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DstIo *string `json:"dst_io,omitempty"`

	// **参数解释**：  已迁移的数据量，用于评估迁移进度。单位：MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	TransData *string `json:"trans_data,omitempty"`

	// **参数解释**：  已迁移的数据行数，用于评估迁移进度。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	TransLines *string `json:"trans_lines,omitempty"`

	// **参数解释**：  迁移实例磁盘使用量，用于监控迁移实例存储容量。单位：GB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	UsedVolumes *string `json:"used_volumes,omitempty"`

	// **参数解释**：  迁移实例内存使用量，用于监控迁移实例内存消耗。单位：MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	UsedMemory *string `json:"used_memory,omitempty"`

	// **参数解释**：  迁移实例CPU使用百分比，用于监控迁移实例计算负载。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	UsedCpuPercent *string `json:"used_cpu_percent,omitempty"`

	// **参数解释**：  迁移实例磁盘总大小，与used_volumes配合评估磁盘使用率。单位：GB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	NodeVolumeSize *int32 `json:"node_volume_size,omitempty"`

	// **参数解释**：  迁移实例内存总大小，与used_memory配合评估内存使用率。单位：MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	NodeMemorySize *int32 `json:"node_memory_size,omitempty"`

	// **参数解释**：  监控数据的更新时间，UTC格式，如2023-07-19T03:09:25Z。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**：  同步速度，用于评估数据回放速率。单位：byte/s。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ApplyRate      *int32 `json:"apply_rate,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMonitorDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorDataResponse struct{}"
	}

	return strings.Join([]string{"ShowMonitorDataResponse", string(data)}, " ")
}
