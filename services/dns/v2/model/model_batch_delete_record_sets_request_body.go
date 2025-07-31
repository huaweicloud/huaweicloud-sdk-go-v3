package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRecordSetsRequestBody struct {

	// **参数解释：** 域名的类型。 **约束限制：** 不涉及。 **取值范围：** - public：公网域名 - private：内网域名  **默认取值：** 不涉及。
	ZoneType string `json:"zone_type"`

	// **参数解释：** 待删除的记录集ID列表。 **约束限制：** 最多支持100个。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RecordsetIds []string `json:"recordset_ids"`
}

func (o BatchDeleteRecordSetsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetsRequestBody", string(data)}, " ")
}
