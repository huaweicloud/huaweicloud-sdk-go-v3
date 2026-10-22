package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentInfoDocCountItems struct {

	// **参数解释：** 统计维度标识。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 该维度的文档计数。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	DocCount *int64 `json:"doc_count,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoDocCountItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoDocCountItems struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoDocCountItems", string(data)}, " ")
}
