package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchCreateRecordSetsTaskRequest Request Object
type ShowBatchCreateRecordSetsTaskRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`

	// **参数解释：** 分页查询时配置每页返回的失败条目个数。 **约束限制：** 不涉及。 **取值范围：** 0~2000。 **默认取值：** 2000
	ErrorItemLimit *int32 `json:"error_item_limit,omitempty"`

	// **参数解释：** 分页查询起始偏移量，表示从偏移量的下一个失败条目开始查询。 **约束限制：** 不涉及。 **取值范围：** 0~2000。 **默认取值：** 0
	ErrorItemOffset *int32 `json:"error_item_offset,omitempty"`
}

func (o ShowBatchCreateRecordSetsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchCreateRecordSetsTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchCreateRecordSetsTaskRequest", string(data)}, " ")
}
