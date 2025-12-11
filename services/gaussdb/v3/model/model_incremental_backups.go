package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncrementalBackups struct {

	// 备份ID。
	Id *string `json:"id,omitempty"`

	// 备份名称。
	Name *string `json:"name,omitempty"`

	// 备份开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	BeginTime *string `json:"begin_time,omitempty"`

	// 备份结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	EndTime *string `json:"end_time,omitempty"`

	// 备份大小，(单位：KB)。
	Size float32 `json:"size,omitempty"`

	// **参数解释**：  增量备份类型。  **取值范围**：  - Log: 同区域增量备份。 - OffSiteLog：跨区域增量备份。
	BackupType *string `json:"backup_type,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o IncrementalBackups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncrementalBackups struct{}"
	}

	return strings.Join([]string{"IncrementalBackups", string(data)}, " ")
}
