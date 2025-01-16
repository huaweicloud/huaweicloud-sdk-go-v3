package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDataSyncResponse Response Object
type ModifyDataSyncResponse struct {

	// **参数解释**：  TaurusDB数据库名。  **取值范围**：  长度限制3~1024位，仅支持英文大小写字母、数字以及下划线。
	SourceDatabaseName *string `json:"source_database_name,omitempty"`

	// **参数解释**：  目标数据库名。  **取值范围**：  长度限制3~1024位，仅支持英文大小写字母、数字以及下划线。
	TargetDatabaseName *string `json:"target_database_name,omitempty"`

	// **参数解释**：  TaurusDB数据库配置检查结果。  **取值范围**：  不涉及。
	SourceDbConfigCheckResults *[]DbConfigCheckResult `json:"source_db_config_check_results,omitempty"`

	// **参数解释**：  表配置检查结果。  **取值范围**：  不涉及。
	TblConfigCheckResults *[]TableConfigCheckResult `json:"tbl_config_check_results,omitempty"`

	// **参数解释**：  同步任务名称。  **取值范围**：  长度限制3~128位，仅支持英文大小写字母、数字以及下划线。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**：  修改同步配置工作流ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyDataSyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDataSyncResponse struct{}"
	}

	return strings.Join([]string{"ModifyDataSyncResponse", string(data)}, " ")
}
