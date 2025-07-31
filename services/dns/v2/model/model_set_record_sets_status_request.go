package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRecordSetsStatusRequest Request Object
type SetRecordSetsStatusRequest struct {

	// **参数解释：** 记录集ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RecordsetId string `json:"recordset_id"`

	Body *SetRecordSetsStatusRequestBody `json:"body,omitempty"`
}

func (o SetRecordSetsStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecordSetsStatusRequest struct{}"
	}

	return strings.Join([]string{"SetRecordSetsStatusRequest", string(data)}, " ")
}
