package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtsConfig **参数解释**：LTS（云日志服务）配置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type LtsConfig struct {

	// **参数解释**：是否开启LTS日志。 **约束限制**：不涉及。 **取值范围**：   - true：开启。   - false：不开启。 **默认取值**：false。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**：LTS日志组ID，可在云日志服务LTS控制台日志管理页面的日志组列表查看。 **约束限制**：enable为true的时候需要配置。 **取值范围**：长度为1~36个字符，支持英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	LtsGroupId *string `json:"lts_group_id,omitempty"`

	// **参数解释**：LTS日志流ID，可在云日志服务LTS控制台日志管理页面的日志组列表展开查看对应的日志流。 **约束限制**：enable为true的时候需要配置。 **取值范围**：长度为1~36个字符，支持英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	LtsStreamId *string `json:"lts_stream_id,omitempty"`
}

func (o LtsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsConfig struct{}"
	}

	return strings.Join([]string{"LtsConfig", string(data)}, " ")
}
