package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRecordSetsResponse Response Object
type BatchDeleteRecordSetsResponse struct {

	// **参数解释：** 记录集列表信息。 **取值范围：** 不涉及。
	Recordsets *[]RecordsetData `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteRecordSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetsResponse", string(data)}, " ")
}
