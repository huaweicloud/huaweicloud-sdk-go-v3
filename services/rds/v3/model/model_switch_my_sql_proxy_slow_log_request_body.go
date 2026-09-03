package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchMySqlProxySlowLogRequestBody 更改数据库代理慢日志上报开关请求体
type SwitchMySqlProxySlowLogRequestBody struct {

	// **参数解释**：  慢日志上报开关。  **约束限制**：  不涉及。  **取值范围**：  - on：开启。 - off：关闭。  **默认取值**：  不涉及。
	LtsSlowLogEnabled string `json:"lts_slow_log_enabled"`
}

func (o SwitchMySqlProxySlowLogRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchMySqlProxySlowLogRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchMySqlProxySlowLogRequestBody", string(data)}, " ")
}
