package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchGaussMySqlProxyBinlogPullRequestBody struct {

	// **参数解释**：  数据库代理模式。  **约束限制**：  不涉及。  **取值范围**： - readwrite：读写模式。 - readonly：TP只读模式。  **默认取值**：  不涉及。
	Type string `json:"type"`

	// **参数解释**：  数据库代理拉取Binlog开关。  **约束限制**：  不涉及。  **取值范围**： - true：开启。 - false：关闭。  **默认取值**：  不涉及。
	SupportBinlogPull bool `json:"support_binlog_pull"`

	// **参数解释**： 数据库代理是否从数据库主节点拉取Binlog。  **约束限制**：  - 读写代理可以设置从数据库主节点或者只读节点拉取Binlog，TP只读代理只支持从只读节点拉取Binlog。 - binlog_route_master参数仅在support_binlog_pull为true时生效。  **取值范围**： - true：从主节点拉取。 - false：从只读节点拉取。  **默认取值**：  false。
	BinlogRouteMaster *bool `json:"binlog_route_master,omitempty"`
}

func (o SwitchGaussMySqlProxyBinlogPullRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxyBinlogPullRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxyBinlogPullRequestBody", string(data)}, " ")
}
