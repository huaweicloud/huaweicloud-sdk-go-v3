package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateRecordSetWithBatchLinesResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Recordsets *[]QueryRecordSetWithLineAndTagsResp `json:"recordsets,omitempty"`

	Metadata       *Metedata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateRecordSetWithBatchLinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithBatchLinesResponse struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithBatchLinesResponse", string(data)}, " ")
}
