package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PackLogInfo binlog合并下载文件信息
type PackLogInfo struct {

	// **参数解释**：  文件唯一ID标识。  **约束限制**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  实例id。  **约束限制**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  文件大小。  **约束限制**：  不涉及。
	Size *float64 `json:"size,omitempty"`

	// **参数解释**：  文件大小单位。  **约束限制**：  不涉及。
	SizeUnit *string `json:"size_unit,omitempty"`

	// **参数解释**：  状态。  **约束限制**：  不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  合并时间段起始时间戳。  **约束限制**：  不涉及。
	QueryStartTime *int64 `json:"query_start_time,omitempty"`

	// **参数解释**：  合并时间段结束时间戳。  **约束限制**：  不涉及。
	QueryEndTime *int64 `json:"query_end_time,omitempty"`

	// **参数解释**：  文件名。  **约束限制**：  不涉及。
	FileName *string `json:"file_name,omitempty"`
}

func (o PackLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackLogInfo struct{}"
	}

	return strings.Join([]string{"PackLogInfo", string(data)}, " ")
}
