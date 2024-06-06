package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProxyIpgroupResponse Response Object
type ShowProxyIpgroupResponse struct {

	// 允许访问控制或者不允许 true | false。
	EnableIpGroup *bool `json:"enable_ip_group,omitempty"`

	// 白名单或者黑名单 'white' | 'black'
	Type *string `json:"type,omitempty"`

	IpGroup        *ProxyIpGroupDetail `json:"ip_group,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowProxyIpgroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProxyIpgroupResponse struct{}"
	}

	return strings.Join([]string{"ShowProxyIpgroupResponse", string(data)}, " ")
}
