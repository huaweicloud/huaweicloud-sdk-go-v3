package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRecordSetRequestBody struct {

	// **参数解释：** 域名，后缀需以zone name结束且为FQDN（Fully Qualified Domain Name，全称域名），即以“.”结束的完整主机名。 如“www.example.com.”。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 记录集的描述信息。 **约束限制：** 不涉及。 **取值范围：** 长度不超过255个字符。 **默认取值：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 记录集的类型。 **约束限制：** 不涉及。 **取值范围：** - A：将域名解析到指定的IPv4地址。 - AAAA：将域名解析到指定的IPv6地址。 - MX：指定域名对应的邮件服务器。 - CNAME：将域名解析到另一域名，或者多个域名映射到同一域名上。 - TXT：用于对域名进行标识和说明。 - SRV：用于记录某台服务器对外提供了哪些服务。 - NS：指定域名的权威DNS服务器。 - SOA：提供域名的基本信息和权威服务器的详细信息。 - CAA：指定为域名颁发HTTPS证书的授权CA机构，仅支持公网域名解析。 - PTR：指定IP地址反向解析记录，仅支持内网域名解析。  **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 解析记录的状态。 **约束限制：** 不涉及。 **取值范围：** - ENABLE：启用解析 - DISABLE：暂停解析  **默认取值：** ENABLE。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。 如果您的服务地址经常更换，建议TTL值设置相对小些，反之，建议设置相对大些。 **约束限制：** 不涉及。 **取值范围：** 1~2147483647。 **默认取值：** 300
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 解析记录的值。不同类型解析记录对应的值的规则不同。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Records []string `json:"records"`

	// **参数解释：** 资源标签。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateRecordSetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRecordSetRequestBody", string(data)}, " ")
}
