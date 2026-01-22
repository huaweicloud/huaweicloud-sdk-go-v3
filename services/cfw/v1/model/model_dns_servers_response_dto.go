package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DnsServersResponseDto struct {

	// **参数解释**： 域名服务器id **取值范围**： 不涉及
	Id *int32 `json:"id,omitempty"`

	// **参数解释**： 域名服务器是否应用 **取值范围**： - 0：否 - 1：是
	IsApplied *int32 `json:"is_applied,omitempty"`

	// **参数解释**： 域名服务器是否是用户自定义的dns服务器 **取值范围**： - 0：否 - 1：是
	IsCustomized *int32 `json:"is_customized,omitempty"`

	// **参数解释**： DNS服务器IP **取值范围**： 不涉及
	ServerIp *string `json:"server_ip,omitempty"`

	// **参数解释**： dns服务器解析状态 **取值范围**： 0：解析域名的频率正常 1：解析域名的频率缓慢 2：解析域名异常
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 健康检查域名 **取值范围**： 不涉及
	HealthCheckDomainName *string `json:"health_check_domain_name,omitempty"`
}

func (o DnsServersResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsServersResponseDto struct{}"
	}

	return strings.Join([]string{"DnsServersResponseDto", string(data)}, " ")
}
