package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProxyFlavorsByAzCodeResponse Response Object
type ShowProxyFlavorsByAzCodeResponse struct {

	// **参数解释**：  代理规格分组信息。
	ProxyFlavorGroups *[]ProxyFlavorGroup `json:"proxy_flavor_groups,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ShowProxyFlavorsByAzCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProxyFlavorsByAzCodeResponse struct{}"
	}

	return strings.Join([]string{"ShowProxyFlavorsByAzCodeResponse", string(data)}, " ")
}
