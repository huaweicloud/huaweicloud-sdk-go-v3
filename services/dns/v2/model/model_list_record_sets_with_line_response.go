package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRecordSetsWithLineResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Recordsets *[]QueryRecordSetWithLineAndTagsResp `json:"recordsets,omitempty"`

	Metadata       *Metedata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsWithLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsWithLineResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsWithLineResponse", string(data)}, " ")
}
