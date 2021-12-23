package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群重启操作，定义哪些集群节点需要重启，请参见restart参数说明
type CdmRestartClusterReqRestart struct {
	// 重启时延，单位：秒

	RestartDelayTime *int32 `json:"restartDelayTime,omitempty"`
	// 重启类型： - IMMEDIATELY：立即重启。 - GRACEFULL：优雅重启。 - FORCELY：强制重启。 - SOFTLY：一般重启。  默认值为“IMMEDIATELY”

	RestartMode *string `json:"restartMode,omitempty"`
	// 重启级别： - SERVICE：重启服务。 - VM：重启虚拟机。  默认值为“SERVICE”。

	RestartLevel *string `json:"restartLevel,omitempty"`
	// 集群节点类型，只支持“cdm”

	Type string `json:"type"`
	// 预留字段，“restartLevel” 为“SERVICE”时，“instance”必填，填空字串。

	Instance *string `json:"instance,omitempty"`
	// 预留字段，“restartLevel” 为“SERVICE”时，“group”必填，填空字串。

	Group *string `json:"group,omitempty"`
}

func (o CdmRestartClusterReqRestart) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmRestartClusterReqRestart struct{}"
	}

	return strings.Join([]string{"CdmRestartClusterReqRestart", string(data)}, " ")
}
