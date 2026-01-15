package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainDetectionRequest Request Object
type ShowDomainDetectionRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`

	// **参数解释：** 待诊断记录集的类型。 **取值范围：** - MX：指定域名对应的邮件服务器。 - CNAME：将域名解析到另一域名，或者多个域名映射到同一域名上。 - TXT：用于对域名进行标识和说明。
	Type *string `json:"type,omitempty"`

	// 待诊断记录集的名称。
	DomainName string `json:"domain_name"`
}

func (o ShowDomainDetectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainDetectionRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainDetectionRequest", string(data)}, " ")
}
