package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyProxyConsistRequest 修改代理一致性请求体
type ModifyProxyConsistRequest struct {

	// 会话一致性。 - 取值\"true\"时表示会话一致性开启。 - 取值\"false\"时表示会话一致性关闭。
	SessionConsistence string `json:"session_consistence"`

	// 一致性模式。默认值为空，此时以会话一致性参数session_consistence为准。 - session: 会话一致性 - global: 全局一致性 - eventual: 最终一致性
	ConsistenceMode *string `json:"consistence_mode,omitempty"`
}

func (o ModifyProxyConsistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyProxyConsistRequest struct{}"
	}

	return strings.Join([]string{"ModifyProxyConsistRequest", string(data)}, " ")
}
