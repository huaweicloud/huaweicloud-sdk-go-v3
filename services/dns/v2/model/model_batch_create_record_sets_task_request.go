package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateRecordSetsTaskRequest Request Object
type BatchCreateRecordSetsTaskRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`

	Body *BatchCreateRecordSetsTaskRequestBody `json:"body,omitempty"`
}

func (o BatchCreateRecordSetsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateRecordSetsTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateRecordSetsTaskRequest", string(data)}, " ")
}
