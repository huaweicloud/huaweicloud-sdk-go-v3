package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlProxyAltRequestBody 更改数据库代理ALT开关的请求体
type SwitchGaussMySqlProxyAltRequestBody struct {

	// **参数解释**：  ALT开关状态。  **取值范围**： - on：开启。 - off：关闭。  **默认取值**： 不涉及。
	AltEnabled string `json:"alt_enabled"`
}

func (o SwitchGaussMySqlProxyAltRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxyAltRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxyAltRequestBody", string(data)}, " ")
}
