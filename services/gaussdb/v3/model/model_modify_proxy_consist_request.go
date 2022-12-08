package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改代理一致性请求体
type ModifyProxyConsistRequest struct {

	// 会话一致性。 - 取值\"true\"时表示会话一致性开启。 - 取值\"false\"时表示会话一致性关闭。
	SessionConsistence string `json:"session_consistence"`
}

func (o ModifyProxyConsistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyProxyConsistRequest struct{}"
	}

	return strings.Join([]string{"ModifyProxyConsistRequest", string(data)}, " ")
}
