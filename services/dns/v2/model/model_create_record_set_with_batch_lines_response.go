package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateRecordSetWithBatchLinesResponse struct {
	Links *PageLink `json:"links,omitempty" xml:"links"`

	Recordsets *[]QueryRecordSetWithLineResp `json:"recordsets,omitempty" xml:"recordsets"`

	Metadata       *Metedata `json:"metadata,omitempty" xml:"metadata"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateRecordSetWithBatchLinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithBatchLinesResponse struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithBatchLinesResponse", string(data)}, " ")
}
