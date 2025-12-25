package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecGuardTaskCount struct {

	// **参数解释**: 扫描次数。 **取值范围**: 不涉及。
	Task *int32 `json:"task,omitempty"`

	// **参数解释**: 扫描制品数。 **取值范围**: 不涉及。
	Artifact *int32 `json:"artifact,omitempty"`

	Opensource *OpensourceCount `json:"opensource,omitempty"`

	// **参数解释**: 病毒文件数。 **取值范围**: 不涉及。
	Virus *int32 `json:"virus,omitempty"`

	Malware *MalwareCount `json:"malware,omitempty"`

	// **参数解释**: 扫描总数。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 扫描任务列表。 **取值范围**: 不涉及。
	List *[]SecGuardTaskDetail `json:"list,omitempty"`
}

func (o SecGuardTaskCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecGuardTaskCount struct{}"
	}

	return strings.Join([]string{"SecGuardTaskCount", string(data)}, " ")
}
