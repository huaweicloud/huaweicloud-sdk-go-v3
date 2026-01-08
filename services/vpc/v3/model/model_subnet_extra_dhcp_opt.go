package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubnetExtraDhcpOpt struct {

	// **参数解释**： 子网配置的NTP地址、DNS域名或租约到期时间的名称。 **取值范围**： - 42：表示子网ntp地址。 - 15：表示DNS配置的域名，用于向DNS服务器获取IP地址。 - 51：表示IPv4子网租约到期时间。 - 651：表示IPv6子网租约到期时间。
	OptName string `json:"opt_name"`

	// **参数解释**： 子网配置的NTP地址、DNS域名或租约到期时间。 **取值范围**： - opt_name配置为“42”，则表示是子网ntp地址，目前只支持IPv4地址，每个IP地址以逗号隔开，IP地址个数不能超过4个，不能存在相同地址。 - opt_name配置为“15”，则该值表示是DNS配置的域名，用于向DNS服务器获取IP地址。域名只能由字母，数字，中划线组成，中划线不能在开头或末尾。域名可以包含多个字符串，单个字符串不超过63个字符，字符串间以点分隔。域名长度不超过254个字符。 - opt_name配置为“51”，则该值表示是IPv4子网租约到期时间，取值格式有两种，取-1，表示无限租约；数字+h，数字范围是1~175200，比如5h，默认值为87600h。 - opt_name配置为“651”，则该值表示是IPv6子网租约到期时间，取值格式有两种，取-1，表示无限租约；数字+h，数字范围是1~175200，比如5h，默认值为2h。
	OptValue string `json:"opt_value"`

	// **参数解释**： 子网的IP版本。 **取值范围**： - 4：默认值，IPv4。 - 6：IPv6，只有opt_name是“651”时，ip_version为6。
	IpVersion int32 `json:"ip_version"`
}

func (o SubnetExtraDhcpOpt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetExtraDhcpOpt struct{}"
	}

	return strings.Join([]string{"SubnetExtraDhcpOpt", string(data)}, " ")
}
