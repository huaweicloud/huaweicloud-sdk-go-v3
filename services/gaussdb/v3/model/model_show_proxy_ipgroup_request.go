package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProxyIpgroupRequest Request Object
type ShowProxyIpgroupRequest struct {

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 数据库代理ID，严格匹配UUID规则。
	ProxyId string `json:"proxy_id"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowProxyIpgroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProxyIpgroupRequest struct{}"
	}

	return strings.Join([]string{"ShowProxyIpgroupRequest", string(data)}, " ")
}
