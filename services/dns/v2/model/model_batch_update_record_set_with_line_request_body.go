package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateRecordSetWithLineRequestBody struct {

	// **参数解释：** 记录集列表。 **约束限制：**  最多支持50个。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Recordsets []BatchUpdateRecordSet `json:"recordsets"`
}

func (o BatchUpdateRecordSetWithLineRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRecordSetWithLineRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateRecordSetWithLineRequestBody", string(data)}, " ")
}
