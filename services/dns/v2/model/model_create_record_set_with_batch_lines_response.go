package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecordSetWithBatchLinesResponse Response Object
type CreateRecordSetWithBatchLinesResponse struct {
	Links *PageLink `json:"links,omitempty"`

	// **参数解释：** 记录集列表信息。 **取值范围：** 不涉及。
	Recordsets *[]QueryRecordSetWithLineResp `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateRecordSetWithBatchLinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithBatchLinesResponse struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithBatchLinesResponse", string(data)}, " ")
}
