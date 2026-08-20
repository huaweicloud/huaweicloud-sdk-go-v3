package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogConfigResponse **参数解释：**  服务日志配置信息。
type LogConfigResponse struct {

	// **参数解释：** 日志输出类型。 **取值范围：** - STDOUT：日志输出到控制台或终端。 - EVENT：k8s事件。 - FILE：容器日志文件
	Type string `json:"type"`

	// **参数解释：** 日志服务状态。 **取值范围：** - ON：开启。 - OFF：关闭。 - FAILED：开启失败。
	Status string `json:"status"`

	// **参数解释：** 日志组ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// **参数解释：** 日志流id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// **参数解释：** 日志流对应的部署ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DeploymentId *string `json:"deployment_id,omitempty"`

	// **参数解释：** 失败原因，部署开启LTS对接失败的原因。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	FailureReason *string `json:"failure_reason,omitempty"`
}

func (o LogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogConfigResponse struct{}"
	}

	return strings.Join([]string{"LogConfigResponse", string(data)}, " ")
}
