package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRecordSetWithLineRequestBody struct {

	// **参数解释：** 记录集ID列表。 **约束限制：** 最多支持100个。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RecordsetIds []string `json:"recordset_ids"`
}

func (o BatchDeleteRecordSetWithLineRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetWithLineRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetWithLineRequestBody", string(data)}, " ")
}
