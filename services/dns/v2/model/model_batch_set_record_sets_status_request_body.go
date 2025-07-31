package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSetRecordSetsStatusRequestBody struct {

	// **参数解释：** 待设置记录集状态。 **约束限制：** 不涉及。 **取值范围：** - DISABLE：暂停解析 - ENABLE：启用解析  **默认取值：** 不涉及。
	Status string `json:"status"`

	// **参数解释：** 待设置记录集ID列表。 **约束限制：** 最多支持50个。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RecordsetIds []string `json:"recordset_ids"`
}

func (o BatchSetRecordSetsStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetRecordSetsStatusRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSetRecordSetsStatusRequestBody", string(data)}, " ")
}
