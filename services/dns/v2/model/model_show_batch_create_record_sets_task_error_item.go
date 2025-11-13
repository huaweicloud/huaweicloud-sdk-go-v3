package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowBatchCreateRecordSetsTaskErrorItem struct {

	// **参数解释：** 域名。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 记录集的类型。 **取值范围：** - 公网域名的记录类型: A、CNAME、MX、AAAA、TXT、SRV、NS、CAA、REDIRECT_URL、FORWARD_URL。 - 内网域名的记录类型: A、CNAME、MX、AAAA、TXT、SRV、PTR。
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
