package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtsInfoOption **参数解释**: LTS对接配置信息。 **约束限制**: 不涉及。
type LtsInfoOption struct {

	// **参数解释**: LTS日志组名称。 **约束限制**: 满足正则匹配表达式校验：^GROUP_GAUSSDB_APS-[a-zA-Z0-9-_]{32}in(14|20)$。 **取值范围**: 通常为GROUP_GAUSSDB_APS-实例ID。 **默认取值**: 不涉及。
	LogGroupName string `json:"log_group_name"`

	// **参数解释**: LTS日志流名称。 **约束限制**: 满足正则匹配表达式校验：^GROUP_GAUSSDB_APS-[a-zA-Z0-9-_]{32}in(14|20)$。 **取值范围**: 通常为STREAM_APS_FULL_SQL-实例ID。 **默认取值**: 不涉及。
	LogStreamName string `json:"log_stream_name"`
}

func (o LtsInfoOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsInfoOption struct{}"
	}

	return strings.Join([]string{"LtsInfoOption", string(data)}, " ")
}
