package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowBatchCreateRecordSetsTaskErrorItem struct {

	// **参数解释：** 域名。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 记录集的类型。 **取值范围：** - A：将域名解析到指定的IPv4地址。 - AAAA：将域名解析到指定的IPv6地址。 - MX：指定域名对应的邮件服务器。 - CNAME：将域名解析到另一域名，或者多个域名映射到同一域名上。 - TXT：用于对域名进行标识和说明。 - SRV：用于记录某台服务器对外提供了哪些服务。 - NS：指定域名的权威DNS服务器，当前接口仅支持公网域名解析。 - REDIRECT_URL：显性URL转发，仅支持公网域名解析。 - FORWARD_URL：隐性URL转发，仅支持公网域名解析。 - CAA：指定为域名颁发HTTPS证书的授权CA机构，仅支持公网域名解析。 - PTR：指定IP地址反向解析记录，仅支持内网域名解析。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 解析线路ID。 **取值范围：** 不涉及。
	Line *string `json:"line,omitempty"`

	// **参数解释：** 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。 **取值范围：** 1~2147483647。
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 解析记录的权重。 **取值范围：** 0~1000。
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释：** 域名解析后的值。 **取值范围：** 不涉及。
	Records *[]string `json:"records,omitempty"`

	// **参数解释：** 记录集状态。 **取值范围：** - ERROR：失败
	Status *string `json:"status,omitempty"`

	// **参数解释：** 错误码。 **取值范围：** 不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释：** 错误信息。 **取值范围：** 不涉及。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ShowBatchCreateRecordSetsTaskErrorItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchCreateRecordSetsTaskErrorItem struct{}"
	}

	return strings.Join([]string{"ShowBatchCreateRecordSetsTaskErrorItem", string(data)}, " ")
}
