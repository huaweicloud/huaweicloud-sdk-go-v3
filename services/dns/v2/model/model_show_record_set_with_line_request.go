package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordSetWithLineRequest Request Object
type ShowRecordSetWithLineRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`

	// **参数解释：** 记录集ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RecordsetId string `json:"recordset_id"`
}

func (o ShowRecordSetWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordSetWithLineRequest", string(data)}, " ")
}
