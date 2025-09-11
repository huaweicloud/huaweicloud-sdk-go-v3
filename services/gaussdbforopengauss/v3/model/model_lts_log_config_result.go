package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsLogConfigResult struct {

	// **参数解释**: 日志类型。 **取值范围**: - audit_log：审计日志。
	LogType *string `json:"log_type,omitempty"`

	// **参数解释**: 关联的LTS日志组ID，若enabled为false，表示最近一次关联的LTS日志组ID。 **取值范围**: 不涉及。
	LtsGroupId *string `json:"lts_group_id,omitempty"`

	// **参数解释**: 关联的LTS日志流ID，若enabled为false，表示最近一次关联的LTS日志流ID。 **取值范围**: 不涉及。
	LtsStreamId *string `json:"lts_stream_id,omitempty"`

	// **参数解释**: 关联的LTS日志流是否启用。 **取值范围**: - true：启用。 - false：关闭。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o LtsLogConfigResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsLogConfigResult struct{}"
	}

	return strings.Join([]string{"LtsLogConfigResult", string(data)}, " ")
}
