package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetRecordSetsStatusResponse Response Object
type BatchSetRecordSetsStatusResponse struct {
	Links *Link `json:"links,omitempty"`

	// **参数解释：** 记录集列表信息。 **取值范围：** 不涉及。
	Recordsets *[]RecordsetData `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchSetRecordSetsStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetRecordSetsStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchSetRecordSetsStatusResponse", string(data)}, " ")
}
