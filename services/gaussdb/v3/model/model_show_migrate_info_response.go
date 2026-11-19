package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMigrateInfoResponse Response Object
type ShowMigrateInfoResponse struct {

	// **参数解释**：  源实例ID。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。
	RdsInstanceId *string `json:"rds_instance_id,omitempty"`

	// **参数解释**：  源实例状态。  **取值范围**：  不涉及。
	RdsStatus *string `json:"rds_status,omitempty"`

	// **参数解释**：  源实例名称。  **取值范围**：  不涉及。
	RdsInstanceName *string `json:"rds_instance_name,omitempty"`

	// **参数解释**：  迁移任务ID。  **取值范围**：  不涉及。
	DrsJobId *string `json:"drs_job_id,omitempty"`

	// **参数解释**：  迁移任务状态。  **取值范围**：  不涉及。
	DrsStatus *string `json:"drs_status,omitempty"`

	// **参数解释**：  源实例读写状态。  **取值范围**：  - ON：读写状态。 - OFF：只读状态。
	RdsReadwriteMode *string `json:"rds_readwrite_mode,omitempty"`

	// **参数解释**：  目标实例读写状态。  **取值范围**：  - ON：读写状态。 - OFF：只读状态。
	TaurusReadwriteMode *string `json:"taurus_readwrite_mode,omitempty"`

	// **参数解释**：G  迁移方向。  **取值范围**：  - up：源实例到TaurusDB实例。 - down：TaurusDB实例到源实例。
	MigrationDirection *string `json:"migration_direction,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o ShowMigrateInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrateInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowMigrateInfoResponse", string(data)}, " ")
}
