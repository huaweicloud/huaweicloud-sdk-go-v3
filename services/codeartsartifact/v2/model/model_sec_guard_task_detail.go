package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecGuardTaskDetail struct {

	// **参数解释**: 任务id。 **取值范围**: 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**: 任务名。 **取值范围**: 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**: 文件扫描路径。 **取值范围**: 不涉及。
	ScanPath *string `json:"scan_path,omitempty"`

	// **参数解释**: 文件名。 **取值范围**: 不涉及。
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: 展示名称。 **取值范围**: 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**: 路径。 **取值范围**: 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释**: 创建时间。 **取值范围**: 不涉及。
	CreatedTime *string `json:"created_time,omitempty"`

	Opensource *OpensourceCount `json:"opensource,omitempty"`

	// **参数解释**: 病毒文件数。 **取值范围**: 不涉及。
	Virus *int32 `json:"virus,omitempty"`

	Malware *MalwareCount `json:"malware,omitempty"`

	// **参数解释**: 状态。 **取值范围**: 不涉及。
	Status *string `json:"status,omitempty"`
}

func (o SecGuardTaskDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecGuardTaskDetail struct{}"
	}

	return strings.Join([]string{"SecGuardTaskDetail", string(data)}, " ")
}
