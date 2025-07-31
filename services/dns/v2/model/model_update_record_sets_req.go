package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRecordSetsReq struct {

	// **参数解释：** 域名，后缀需以zone name结束且为FQDN（Fully Qualified Domain Name，全称域名），即以“.”结束的完整主机名。 如“www.example.com.”。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 记录集的描述信息。 **约束限制：** 不涉及。 **取值范围：** 长度不超过255个字符。 **默认取值：** 默认为空，表示维持原值。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 记录集的类型。 **约束限制：** 不涉及。 **取值范围：** - 公网域名的记录类型: A、AAAA、MX、CNAME、TXT、SRV、NS、SOA、CAA。 - 内网域名的记录类型: A、AAAA、MX、CNAME、TXT、PTR、SRV、NS、SOA。  **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。 **约束限制：** 不涉及。 **取值范围：** 1~2147483647。 **默认取值：** 默认为空，表示维持原值。
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 解析记录的值。不同类型解析记录对应的值的规则不同。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 默认为空，表示维持原值。
	Records *[]string `json:"records,omitempty"`

	// **参数解释：** 解析记录的权重。 **约束限制：** 不涉及。 **取值范围：** 取值范围：0~1000。 - 当weight为空时，表示该解析记录将保持原有设置的权重。 - 当weight=0，表示该解析记录为备用域名解析记录。 - 当weight>0，表示该解析记录为主用域名解析记录。  **默认取值：** 默认为空，表示维持原值。
	Weight *int32 `json:"weight,omitempty"`
}

func (o UpdateRecordSetsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordSetsReq struct{}"
	}

	return strings.Join([]string{"UpdateRecordSetsReq", string(data)}, " ")
}
