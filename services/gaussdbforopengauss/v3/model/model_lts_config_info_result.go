package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtsConfigInfoResult **参数解释**: LTS对接配置信息。 **取值范围**: 不涉及。
type LtsConfigInfoResult struct {

	// **参数解释**: LTS日志组名称。 **取值范围**: 不涉及。
	LogGroupName *string `json:"log_group_name,omitempty"`

	// **参数解释**: LTS日志组ID。 **取值范围**: 不涉及。LTS日志组ID
	LogGroupId *string `json:"log_group_id,omitempty"`

	// **参数解释**: LTS日志组类别。 **取值范围**: 通常为asp_log，标识为智能运维专用日志组。
	GroupLogType *string `json:"group_log_type,omitempty"`

	// **参数解释**: LTS日志组中数据最大保留天数。 **取值范围**: [1,30]
	GroupTtlInDays *int32 `json:"group_ttl_in_days,omitempty"`

	// **参数解释**: LTS日志流名称。 **取值范围**: 通常为STREAM_APS_FULL_SQL-实例ID。
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// **参数解释**: LTS日志流ID。 **取值范围**: 不涉及。
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// **参数解释**: LTS日志流类别。 **取值范围**: 通常为full_sql，标识为全量SQL专用日志流。
	StreamLogType *string `json:"stream_log_type,omitempty"`

	// **参数解释**: LTS日志流中数据最大保留天数。 **取值范围**: [1,30]
	StreamTtlInDays *int32 `json:"stream_ttl_in_days,omitempty"`

	// **参数解释**: LTS日志流结构化配置ID。 **取值范围**: 不涉及。
	StreamStructureConfigId *string `json:"stream_structure_config_id,omitempty"`

	// **参数解释**: LTS日志流索引配置ID。 **取值范围**: 不涉及。
	StreamIndexConfigId *string `json:"stream_index_config_id,omitempty"`
}

func (o LtsConfigInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsConfigInfoResult struct{}"
	}

	return strings.Join([]string{"LtsConfigInfoResult", string(data)}, " ")
}
