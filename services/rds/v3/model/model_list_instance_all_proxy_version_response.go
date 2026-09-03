package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceAllProxyVersionResponse Response Object
type ListInstanceAllProxyVersionResponse struct {

	// **参数解释**：  数据库代理节点引擎版本信息列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ProxyEngineVersionInfos *[]ProxyEngineVersionInfo `json:"proxy_engine_version_infos,omitempty"`
	HttpStatusCode          int                       `json:"-"`
}

func (o ListInstanceAllProxyVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceAllProxyVersionResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceAllProxyVersionResponse", string(data)}, " ")
}
