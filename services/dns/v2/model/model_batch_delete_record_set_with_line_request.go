package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRecordSetWithLineRequest Request Object
type BatchDeleteRecordSetWithLineRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`

	Body *BatchDeleteRecordSetWithLineRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteRecordSetWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetWithLineRequest", string(data)}, " ")
}
