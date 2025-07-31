package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRSetBatchLinesReq struct {

	// **参数解释：** 域名，后缀需以zone name结束且为FQDN（Fully Qualified Domain Name，全称域名），即以“.”结束的完整主机名。 如“www.example.com.”。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 记录集的描述信息。 **约束限制：** 不涉及。 **取值范围：** 长度不超过255个字符。 **默认取值：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 记录集的类型。 **约束限制：** 不涉及。 **取值范围：** A、AAAA、MX、CNAME、TXT、SRV、NS、SOA、CAA。 **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 解析线路域名参数。 **约束限制：** 最多支持50个。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Lines []BatchCreateRecordSetWithLine `json:"lines"`
}

func (o CreateRSetBatchLinesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRSetBatchLinesReq struct{}"
	}

	return strings.Join([]string{"CreateRSetBatchLinesReq", string(data)}, " ")
}
