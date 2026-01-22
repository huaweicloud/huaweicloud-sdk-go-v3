package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigQuotaRequest Request Object
type ShowConfigQuotaRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 防火墙配额类型 **约束限制**： 不涉及 **取值范围**：   ACL：ACL规则配额   DNS_DOMAIN_SET：网络域名组配额   DOMAIN：域名组域名成员配额   DOMAIN_DEVICE：域名设备配额   DNS_DOMAIN_SET_PARSE_IP：网络域名组解析IP配额   APPLICATION_DOMAIN_SET：应用域名组配额   APPLICATION_DOMAIN_SET_ITEM：应用域名组域名成员配额   APPLICATION_DOMAIN_SET_ITEM_DEVICE：应用域名组设备配额   ADDR_SET：地址组配额   ADDR_SET_ITEM：地址组IP地址成员配额   ADDR_SET_ITEM_DEVICE：地址组IP地址设备配额   SERV_SET：服务组配额   SERV_SET_ITEM：服务组服务成员配额   SERV_SET_ITEM_DEVICE：服务组服务设备配额   BLACKLIST：黑名单配额   WHITELIST：白名单配额   PRIVATE_NETWORK_SEGMENT：私网网段配额 **默认取值**： 不涉及
	ConfigType string `json:"config_type"`

	// **参数解释**： 组Id，查询IP地址组地址成员配额，域名组域名成员配额，服务组服务成员配额时必传，地址组id，可通过查询地址组列表接口查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得，域名组id，可通过查询域名组列表接口查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得，服务组id，可通过获取服务组列表接口查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	SetId *string `json:"set_id,omitempty"`
}

func (o ShowConfigQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigQuotaRequest", string(data)}, " ")
}
