package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDataSyncConfigRequestV3 StarRocks修改配置请求体。
type ModifyDataSyncConfigRequestV3 struct {

	// **参数解释**：  TaurusDB实例ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认值**：  不涉及。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// **参数解释**：  TaurusDB只读节点ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为no07，长度为36个字符。  **默认值**：  不涉及。
	SourceNodeId *string `json:"source_node_id,omitempty"`

	// **参数解释**：  TaurusDB数据库名。  **约束限制**：  不涉及。  **取值范围**：  字符长度限制3~1024位，仅支持英文大小写字母、数字以及下划线。  **默认值**：  不涉及。
	SourceDatabaseName *string `json:"source_database_name,omitempty"`

	// **参数解释**：  数据同步任务名称。  **约束限制**：  不涉及。  **取值范围**：  长度限制3~128位，仅支持英文大小写字母、数字以及下划线。  **默认值**：  不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**：  库配置列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认值**：  不涉及。
	DbConfigs *[]DbConfig `json:"db_configs,omitempty"`

	// **参数解释**：  表配置信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认值**：  不涉及。
	TablesConfigs *[]TablesConfig `json:"tables_configs,omitempty"`

	TableReplConfig *TableReplConfig `json:"table_repl_config,omitempty"`

	// **参数解释**：  目标数据库名。  **约束限制**：  不涉及。  **取值范围**：  长度限制3~128位，仅支持英文大小写字母、数字以及下划线。  **默认值**：  不涉及。
	TargetDatabaseName *string `json:"target_database_name,omitempty"`

	// **参数解释**：  是否支持实例级同步。  **约束限制**：  不涉及。  **取值范围**：  - true：是。 - false：否。  **默认取值**：  false。
	IsInstanceLevelSync *string `json:"is_instance_level_sync,omitempty"`

	// **参数解释**：  库同步范围。  **约束限制**：  不涉及。  **取值范围**：  - all：同步全部库。 - part：同步部分库。  **默认取值**：  part。
	DatabaseReplScope *string `json:"database_repl_scope,omitempty"`

	// **参数解释**：  是否支持通配符。  **约束限制**：  不涉及。  **取值范围**：  - true：支持通配符。 - false：不支持通配符。  **默认取值**：  false。
	IsSupportRegExp *string `json:"is_support_reg_exp,omitempty"`
}

func (o ModifyDataSyncConfigRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDataSyncConfigRequestV3 struct{}"
	}

	return strings.Join([]string{"ModifyDataSyncConfigRequestV3", string(data)}, " ")
}
