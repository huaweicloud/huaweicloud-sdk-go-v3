package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtsConfig **参数解释：** 日志配置。 **约束限制：** 不涉及。
type LtsConfig struct {

	// **参数解释：** 日志服务状态。 **取值范围：** - ON：开启。 - OFF：关闭。
	Status string `json:"status"`

	// **参数解释：** LTS日志类型。 **约束限制：** 不涉及。 **取值范围：** - STDOUT：标准日志输入输出 - EVENT：Kubernetes事件 **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 日志组ID，用户选择自己已有的日志组，不填时，会自动创建。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// **参数解释：** 日志流id，用户选择自己已有的日志组。不填时，会自动创建。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	LogStreamId *string `json:"log_stream_id,omitempty"`
}

func (o LtsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsConfig struct{}"
	}

	return strings.Join([]string{"LtsConfig", string(data)}, " ")
}
