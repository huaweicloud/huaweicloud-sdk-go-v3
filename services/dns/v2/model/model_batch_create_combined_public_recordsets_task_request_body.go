package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateCombinedPublicRecordsetsTaskRequestBody struct {

	// **参数解释：** 托管该记录的域名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneNames []string `json:"zone_names"`

	// **参数解释：** 主机记录。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RecordsetNamePrefixes []string `json:"recordset_name_prefixes"`

	// **参数解释：** 解析记录的值。不同类型解析记录对应的值的规则不同。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Records []string `json:"records"`

	// **参数解释：** 记录集的类型。 **约束限制：** 不涉及。 **取值范围：** - 公网域名的记录类型: A、CNAME、MX、AAAA、TXT、SRV、NS、CAA。  **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 解析线路ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** default_view。
	Line *string `json:"line,omitempty"`

	// **参数解释：** 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。 **约束限制：** 不涉及。 **取值范围：** 1~2147483647。 **默认取值：** 300
	Ttl *int32 `json:"ttl,omitempty"`
}

func (o BatchCreateCombinedPublicRecordsetsTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateCombinedPublicRecordsetsTaskRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateCombinedPublicRecordsetsTaskRequestBody", string(data)}, " ")
}
