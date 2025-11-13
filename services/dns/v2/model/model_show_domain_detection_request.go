package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainDetectionRequest Request Object
type ShowDomainDetectionRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`

	// 待诊断记录集的类型。 取值范围：CNAME、TXT、MX。
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
